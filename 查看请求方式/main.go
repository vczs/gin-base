package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Any 根据请求方式参数进行对应处理
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet: //GET请求
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost: //POST请求
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		case http.MethodPut: //PUT请求
			c.JSON(http.StatusOK, gin.H{"method": "PUT"})
			//...
		}
	})
}
