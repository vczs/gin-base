package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//url自定义参数 url中前面有":"的都是参数  ":"后是参数名  可以在gin.Context中通过c.Param()传入参数名获取
	//定义任何参数都不能为空  只要有任何参数为空 就返回404
	r.GET("/user/:admin/:password", func(c *gin.Context) {
		admin := c.Param("admin")
		password := c.Param("password")
		c.JSON(http.StatusOK, gin.H{"admin": admin, "password": password})
	})
}
