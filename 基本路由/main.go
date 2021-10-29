package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//gin框架
func main() {
	r := gin.Default()                   //返回默认的路由引擎
	r.LoadHTMLFiles("./view/model.tmpl") //解析模板
	//指定用户使用GET请求访问/hello时 执行r.GET()函数的第二个参数该匿名函数 渲染模板
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "c1.tmpl", gin.H{"title": "您好"}) })
	r.Run(":9001") //在9001端口 启动服务
}
