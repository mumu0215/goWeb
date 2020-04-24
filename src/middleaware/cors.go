package middleaware

import "github.com/gin-gonic/gin"

func Cors()gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin","*")
		context.Next()
	}
}
