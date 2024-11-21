package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html 的话，
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
