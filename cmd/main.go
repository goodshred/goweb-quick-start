package main

import (
	"goweb-quick-start/api/router"
	"goweb-quick-start/dao"
)

func init() {
	dao.InitMysql()
}

func main() {
	/**
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	*/

	router.Run()
}
