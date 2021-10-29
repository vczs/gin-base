package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由引擎
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx", "./statics")
	//自定义模板要使用的方法
	r.SetFuncMap(template.FuncMap{"safe": func(url string) template.HTML { return template.HTML(url) }})
	//解析模板
	//r.LoadHTMLFiles("./views/v1/index.tmpl", "./views/v2/index.tmpl") //指定模板解析
	r.LoadHTMLGlob("./views/**/*") //解析多个模板
	//渲染模板
	r.GET("/c1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index01", gin.H{"title": "<script>alert(123456789);</script>"})
	})
	r.GET("/c2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index02", gin.H{"title": "<a href='https://www.baidu.com'>百度</a>"})
	})
	//启动服务
	r.Run(":9002")
}
