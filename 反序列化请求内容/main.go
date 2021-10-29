package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type People struct {
	Admin    string `form:"admin"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	//将url请求中的值赋值给结构体
	r.GET("/people", func(c *gin.Context) {
		p := &People{}
		err := c.ShouldBind(p) //通过反射将请求中的值一一对应赋值给结构体字段
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": p})
		}
	})

	r.Run(":9004")
}
