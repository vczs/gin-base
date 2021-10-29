package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由引擎
	r := gin.Default()

	//解析模板
	r.LoadHTMLFiles("./views/index.html", "./views/home.html")

	//渲染模板
	r.GET("/h1", func(c *gin.Context) {
		// name := c.Query("query") //获取请求url中的query的值
		name, _ := c.GetQuery("query") //和c.Query()函数相同 多个bool返回值 如果获取不到bool返回值为false
		//name := c.DefaultQuery("query", "abcd") //获取请求url中的query的值 如果url中没有query 就返回abcd给name
		c.JSON(http.StatusOK, gin.H{"name": name})
	})

	r.GET("/h2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/h3", func(c *gin.Context) {
		//获取form表单内容
		//admin := c.PostForm("admin")
		//password := c.PostForm("password")

		//获取form表单内容 如果没有该值 就返回自定义内容
		admin := c.DefaultPostForm("admin", "admin123")
		password := c.DefaultPostForm("password", "password123")
		c.HTML(http.StatusOK, "home.html", gin.H{
			"admin":    admin,
			"password": password,
		})
	})

	r.Run(":9003")
}
