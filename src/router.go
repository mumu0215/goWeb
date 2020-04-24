package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
	. "src/apis"
	"src/common"
	"src/database"
	"src/middleaware"
	"src/models"
	"src/myvalidate"
)
func initRouter()*gin.Engine  {
	// validator v8到v9
	binding.Validator = new(myvalidate.DefaultValidator)

	app:=gin.New()
	if !database.MyDb.HasTable(models.Products{}){
		if err:=database.MyDb.CreateTable(models.Products{}).Error;err!=nil{
			fmt.Println(err)
			os.Exit(1)
		}
	}
	app.Use(middleaware.GenerateLog())

	app.NoRoute(func(context *gin.Context) {
		response:=&common.Response{}
		response.Msg="No such path"
		response.Success=false
		context.JSON(http.StatusBadRequest,response)
		context.Abort()
	})
	app.NoMethod(func(context *gin.Context) {
		response:=&common.Response{}
		response.Msg="No such method"
		response.Success=false
		context.JSON(http.StatusMethodNotAllowed,response)
		context.Abort()
	})
	app.Use(gin.Recovery())
	app.Use(middleaware.Cors())
	pprof.Register(app)

	userGroup :=app.Group("/user")
	{
		userGroup.POST("/login",UserLogin)
		userGroup.GET("/logout",UserLogout)
		userGroup.POST("/register",Register)
	}

	productGroup:=app.Group("/index")
	{
		productGroup.GET("/",IndexApi)
		productGroup.GET("/Products",FindProductApi)
		productGroup.GET("/Product/:id",FindProductApi)
		productGroup.POST("/Product",AddProductApi)
		productGroup.DELETE("/Product/:id",DeleteProductApi)
		productGroup.POST("Products",UpdateProductApi)
	}

	return app
}