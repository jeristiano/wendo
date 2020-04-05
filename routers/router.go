package routers

import (
	"github.com/gin-gonic/gin"
	setting "github.com/jeristiano/wendo/pkg"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "i am well",
		})
	})
	return r
}
