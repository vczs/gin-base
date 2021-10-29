package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//重定向1
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})

	//重定向2
	r.GET("/a", func(c *gin.Context) {
		//跳转到 /b 路由器
		c.Request.URL.Path = "/b" //指定要跳转的url路径
		r.HandleContext(c)        //传入c 执行跳转
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello"})
	})

	r.Run(":9008")
}
