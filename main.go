package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoutes(router)

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
