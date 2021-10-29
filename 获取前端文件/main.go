package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//解析模板
	r.LoadHTMLFiles("./index.html")
	//渲染模板 get请求
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//渲染模板 post请求
	r.POST("/upload", func(c *gin.Context) {
		//获取文件
		f, _ := c.FormFile("fileName")
		//保存文件
		dst := fmt.Sprintf("./%v_%v", time.Now().Unix(), f.Filename)
		err := c.SaveUploadedFile(f, dst)
		if err != nil {
			fmt.Println("保存失败:", err)
		}
	})
	r.Run(":9007")
}
