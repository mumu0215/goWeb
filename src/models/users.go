package models

import(
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	db "src/database"
	"src/middleaware"
	"time"
)
type Users struct {
	ID int `gorm:"primary_key;column:ID"`
	UserName string `gorm:"not null;column:UserName"`
	PassWordHash string `gorm:"not null;column:PassWordHash"`
}

type UsersLog struct {
	ID int `gorm:"primary_key;column:ID"`
	UserName string `gorm:"not null;column:UserName"`
	IP string `gorm:"column:IP"`
	Action string `gorm:"not null;column:action"`
	TimeStamp time.Time `gorm:"not null;type:timestamp;column:timestamp"`
}

func init()  {
	db.MyDb.AutoMigrate(Users{},UsersLog{})
}

//密码验证在这，不在上层实现
//err放在上层处理
func GetHashPassword(name string,password string) (int,error){
	tempUser:=Users{}
	sqlAction:=db.MyDb.Begin()
	defer func() {
		if err:=recover();err!=nil{
			sqlAction.Rollback()
			middleaware.Error(err)
		}
	}()
	err:=sqlAction.Table("Users").Where("UserName=?",name).Find(&tempUser).Error
	if err==gorm.ErrRecordNotFound{
		sqlAction.Rollback()
		return 0,err
	}
	check:=checkPassWordHash([]byte(tempUser.PassWordHash),[]byte(password))
	if !check{                   //待添加redis验证次数限制
		err=sqlAction.Create(&UsersLog{
			UserName:  name,
			Action:    "密码错误",
			TimeStamp: time.Now(),
		}).Error
		if err!=nil{
			sqlAction.Rollback()
			fmt.Println("Cannot save user log infor.")
		}else {sqlAction.Commit()}
		return 0,errors.New("wrong password")
	}
	err=sqlAction.Create(&UsersLog{
		UserName:  name,
		Action:    "登录成功",
		TimeStamp: time.Now(),
	}).Error
	if err!=nil{
		sqlAction.Rollback()
	}else {sqlAction.Commit()}
	return tempUser.ID,err
}

func UserLogout(name string) error {
	sqlAction:=db.MyDb.Begin()
	defer func() {
		if err:=recover();err!=nil{
			sqlAction.Rollback()
			middleaware.Error(err)
		}
	}()
	tempUserLog :=UsersLog{
		UserName:  name,
		Action:    "账户登出",
		TimeStamp: time.Now(),
	}
	err:=sqlAction.Create(&tempUserLog).Error
	if err!=nil{
		sqlAction.Rollback()
		return err
	}else {sqlAction.Commit()}
	return nil
}

func UserRegister(name string,password string) (int,error) {
	sqlAction:=db.MyDb.Begin()
	defer func() {
		if err:=recover();err!=nil{
			sqlAction.Rollback()
			middleaware.Error(err)
		}
	}()

	//校验用户名是否已存在
	// 0:用户存在		-1:未知错误
	tempUser:=Users{}
	err:=sqlAction.Table("users").Where("username=?",name).First(&tempUser).Error
	switch err {
	case gorm.ErrRecordNotFound:
		tempUser.UserName=name
		PassWordHashByte,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err!=nil{
			middleaware.Error(err)
			sqlAction.Rollback()
			return -1,err
		}
		tempUser.PassWordHash=string(PassWordHashByte)
		err=sqlAction.Create(&tempUser).Error
		if err!=nil{
			middleaware.Error(err)
			sqlAction.Rollback()
			return -1,err
		}
		sqlAction.Commit()
		return tempUser.ID,nil
	case nil:
		return 0,nil
	default:           //未知错误
		middleaware.Error(err)
		sqlAction.Rollback()
		return -1,err
	}
}
func checkPassWordHash(hashPass []byte,pass []byte)  bool{
	return nil==bcrypt.CompareHashAndPassword(hashPass,pass)
}
