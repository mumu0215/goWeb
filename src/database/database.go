package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var MyDb *gorm.DB

func init()  {
	var err error
	MyDb,err=gorm.Open("mysql","root:aaron@123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	DealWithErr(err)
	err=MyDb.DB().Ping()
	DealWithErr(err)
}
func DealWithErr(err error)  {
	if err!=nil{
		log.Fatal(err.Error())
	}
}
