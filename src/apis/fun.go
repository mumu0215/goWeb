package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"src/database"
	"src/models"
	"strconv"
	"strings"
)
var errfun error
func IndexApi(content *gin.Context)  {
	content.Writer.WriteString("<h1>hello</h1>")
	//content.
}
func AddProductApi(content *gin.Context)  {
	p:=models.Products{}
	p.ID,errfun=strconv.Atoi(content.PostForm("id"))
	database.DealWithErr(errfun)
	p.Product=content.PostForm("product")
	p.Details=content.PostForm("details")
	p.AddProducts()
	content.JSON(200,"done!")
	//content.Redirect(302,"/")
}
func DeleteProductApi(content *gin.Context)  {
	p:=models.Products{}
	p.ID,errfun=strconv.Atoi(content.Param("id"))
	p.DelProducts()
	content.JSON(200,"done!")
}
func UpdateProductApi(content *gin.Context)  {
	p:=models.Products{}
	t:=models.Products{}
	p.ID,errfun=strconv.Atoi(content.PostForm("id"))
	database.DealWithErr(errfun)
	p.Product=content.PostForm("product")
	p.Details=content.PostForm("details")
	t.UpDateProducts(p)
	content.JSON(200,"done!")
}
func FindProductApi(content *gin.Context)  {
	p:=models.Products{}
	getpath:=content.FullPath()
	fmt.Println(getpath)
	index:=strings.Index(getpath,"Products")
	if index>0{
		content.JSON(200,models.GetAll())
		//content.Redirect(302,"/")
	}else {
		//id:=content.Query("id")
		id:=content.Param("id")
		fmt.Println(id)
		p.ID,errfun=strconv.Atoi(id)
		content.JSON(200,p.FindProducts())
	}
}