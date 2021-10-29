package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//路由组:将公用的前缀提取出来 创建一个路由组
	urlGroup := r.Group("/a")
	{ //这里的大括号不影响代码逻辑 只是为了美观
		//此处完整url是:"/a/aa"
		urlGroup.GET("/aa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "aa"})
		})
		//此处完整url是:"/a/bb"
		urlGroup.GET("/bb", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "bb"})
		})
		//路由组支持嵌套：urlGroup路由组中再嵌套一个urlGroupPlus路由组
		urlGroupPlus := urlGroup.Group("/b")
		{
			//此处完整url是:"/a/b/cc"
			urlGroupPlus.GET("/cc", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"msg": "cc"})
			})
		}
	}

	//设置没有路由的页面 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"404": "没有该页面"})
	})

	r.Run(":9009")
}
