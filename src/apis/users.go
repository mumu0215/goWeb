package apis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	."src/common"
	"src/jwt"
	"src/models"
)

func UserLogin(context *gin.Context)  {
	var userlogin UserLoginForm
	response:=Response{}
	err:=context.ShouldBindJSON(&userlogin)
	if err!=nil{
		response.Success=false
		response.Msg="Unexpect form msg"
		context.JSON(http.StatusBadRequest,response)
	}
	//验证账户密码
	UserID,err:=models.GetHashPassword(userlogin.UserName,userlogin.PassWord)
	switch err {
	case gorm.ErrRecordNotFound:
		response.Success=false
		response.Msg="User not found"
		context.JSON(http.StatusOK,response)
	case nil:
		//验证成功，需要产生token头
		userToken:=jwt.GenerateJWT([]byte(jwt.MySecret),UserToken{
			ID:       UserID,
			UserName: userlogin.UserName,
		})
		context.Header(DefaultJwtTokenHeader,userToken)
		response.Success=true
		response.Msg="Success login"
		response.Data=UserToken{
			ID:       UserID,
			UserName: userlogin.UserName,
		}
		context.JSON(http.StatusOK,response)
	case errors.New("wrong password"):
		response.Success=false
		response.Msg="Wrong password"
		context.JSON(http.StatusOK,response)
	default:
		response.Success=false
		response.Msg="Log failure, Please retry"
		context.JSON(http.StatusInternalServerError,response)
	}
}
func UserLogout(context *gin.Context)  {
	userToken:=context.GetHeader(DefaultJwtTokenHeader)
	tempUserInfo,err:=jwt.ParseToken(userToken)
	response:=Response{}
	responseCode:=http.StatusOK
	if err!=nil{
		response.Msg="Wrong token"
		response.Success=false
		//context.JSON(http.StatusOK,response)
		responseCode=http.StatusOK
	}
	err=models.UserLogout(tempUserInfo.UserName)
	if err!=nil{
		response.Success=false
		//context.JSON(http.StatusInternalServerError,response)
		responseCode=http.StatusInternalServerError
	}else {
		response.Success=true
		response.Msg="Sucess logout"
		//context.JSON(http.StatusOK,response)
		responseCode=http.StatusOK
	}
	context.JSON(responseCode,response)  //最后返回
}
func Register(context *gin.Context)  {
	registerForm:=RegisterForm{}
	responseCode:=http.StatusOK
	response:=Response{}
	err:=context.ShouldBindJSON(&registerForm)
	if err!=nil{
		responseCode=http.StatusBadRequest
		response.Msg="Bad request"
		response.Success=false
	}
	UserId,err:=models.UserRegister(registerForm.UserName,registerForm.PassWord1)
	switch UserId {
	case -1:
		responseCode=http.StatusInternalServerError
		response.Success=false
		response.Msg=err.Error()
	case 0:
		response.Success=false
		responseCode=http.StatusOK
		response.Msg="User already exist"
	default:
		userToken:=jwt.GenerateJWT([]byte(jwt.MySecret),UserToken{
			ID:       UserId,
			UserName: registerForm.UserName,
		})
		context.Header(DefaultJwtTokenHeader,userToken)
		response.Success=true
		responseCode=http.StatusOK
		response.Msg="Register success"
	}
	context.JSON(responseCode,response)
}