package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/a", handler, func(c *gin.Context) {
		val, _ := c.Get("startTime")
		start, _ := val.(time.Time)
		cost := time.Since(start) //从start执行到当前行所用的时间
		fmt.Println("耗时:", cost)
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	r.GET("/b", handler, func(c *gin.Context) {
		val, _ := c.Get("startTime")
		fmt.Println("传递过来的信息:", val)
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	r.Run(":9010")
}

//handler处理器函数
func handler(c *gin.Context) {
	var flag bool = true
	start := time.Now()
	c.Set("startTime", start) //给后面放行的处理器传递数据
	if flag {
		c.Next() //放行后续的处理器
	} else {
		c.Abort() //阻止后续的处理器
	}
}
