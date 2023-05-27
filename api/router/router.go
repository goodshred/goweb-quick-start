package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goweb-quick-start/api/handler"
)

func Run() {
	router := gin.Default()
	//引入go-playground/validator参数校验
	binding.Validator = new(defaultValidator)

	root := router.Group("/api/v1")
	{
		group := root.Group("/student")
		{
			group.GET("", handler.FrontendGetStudent)
			group.POST("", handler.FrontendPostStudent)
			group.DELETE("/:id", handler.FrontendDeleteStudent)
			group.PUT("", handler.FrontendPutStudent)
		}
	}
	router.Run(":8080")
}
