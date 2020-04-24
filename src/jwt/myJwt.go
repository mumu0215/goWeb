package jwt

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"src/common"
	"src/middleaware"

	//"github.com/dgrijalva/jwt-go/request"
)


const MySecret  = "bad_monkey"

func GenerateJWT(secret []byte,info interface{})  string{
	myToken:=jwt.New(jwt.SigningMethodHS256)
	claim:=make(jwt.MapClaims)
	switch t:=info.(type){
	case common.UserToken:
		claim["Id"]=t.ID
		claim["Name"]=t.UserName
	default:
		fmt.Println("Unhandled err!")
	}
	myToken.Claims=claim
	code,err:=myToken.SignedString(secret)
	if err!=nil{fmt.Println("1")}
	return code
}

//解析token的函数keyfunc参数需要接受匿名函数
func ParseToken(token string) (common.UserToken,error) {
	parsedToken,err:=jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		if token.Method!=jwt.SigningMethodHS256{
			return nil,errors.New("bad Token!")
		}
		return []byte(MySecret),nil
	})
	if err!=nil{
		middleaware.Error(err)
		return common.UserToken{},err
	}
	myuser:=common.UserToken{}
	var e error
	claim,ok:=parsedToken.Claims.(jwt.MapClaims)
	if ok&&parsedToken.Valid{
		myuser.ID=int(claim["Id"].(float64))
		myuser.UserName=claim["Name"].(string)
		return myuser,nil
	}else {
		fmt.Println("failed to parse token!")
		e=errors.New("bad parse")
		return common.UserToken{},e
	}
}
