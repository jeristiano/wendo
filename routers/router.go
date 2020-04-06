package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeristiano/wendo/pkg/logging"
	v1 "github.com/jeristiano/wendo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(context *gin.Context) {

		logging.Debug(context.Request.ContentLength)
		context.JSON(200, gin.H{
			"message": "i am well",
		})
	})

	apiv1 := r.Group("/api/vi")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)

		//编辑
		apiv1.PUT("/tags/：id", v1.EditTag)

		//删除
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
