package common

const (
	DefaultJwtTokenHeader="Authorization"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     interface{} `json:"msg"`
}
type UserLoginForm struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}
type UserToken struct {
	ID int `json:"id"`
	UserName string `json:"username"`
}
type RegisterForm struct {
	UserName string `json:"username" form:"username"`
	PassWord1 string `json:"password1" form:"password1"`
	PassWord2 string `json:"password2" form:"password2"`
}