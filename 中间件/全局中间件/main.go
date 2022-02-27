package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(handler1, handler2) //全局注册中间件handler1和handler2
	//以下代码变量r定义的请求 在执行其处理器之前都要先执行注册的全局中间件handler1和handler2
	//执行完注册的全局中间件后 才能执行自己的处理器函数
	r.GET("/", func(c *gin.Context) {
		fmt.Println("执行了/的函数")
		c.JSON(http.StatusOK, gin.H{"success": "hello golang!"})
	})
	r.Run(":9011")
}

//全局注册中间件
func handler1(c *gin.Context) {
	fmt.Println("执行了Handler2函数")
	//c.Abort() //阻止后续的处理器
	fmt.Println("handler2放行了后续处理器")
	c.Next() //执行后续的处理器
	fmt.Println("Handler2函数执行结束")
}
func handler2(c *gin.Context) {
	fmt.Println("执行了Handler3函数")
	//c.Abort() //阻止后续的处理器
	fmt.Println("handler3放行了后续处理器")
	c.Next() //执行后续的处理器
	fmt.Println("Handler3函数执行结束")
}
