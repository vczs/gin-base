package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id     int
	Name   string
	Gender string
	Hobby  string
}

var (
	db  *gorm.DB
	err error
)

func main() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gin_user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c() //创建表
	z() //增
	q() //查
	u() //改
	s() //删
}

func c() {
	//创建表 把结构体字段和数据表一一对应
	db.AutoMigrate(&UserInfo{})
}

func z() {
	//增：向数据表添加一条数据
	u1 := &UserInfo{Id: 1, Name: "王", Gender: "男", Hobby: "双色球"}
	//增加
	db.Create(u1) //添加到数据表
}

func q() {
	//查：从数据表查询一条数据
	var u1 UserInfo
	//查询
	db.First(&u1) //从数据表查询第一条数据保存到u1
	fmt.Printf("%#v\n", u1)
}

func u() {
	//改：从数据表修改一条数据
	var u1 UserInfo
	db.First(&u1) //从数据表查询第一条数据保存到u1
	//修改
	//db.Save(&u1) //更新所有字段
	db.Model(&u1).Update("hobby", "彩票") //只更新指定的字段
	q()                                 //再查询一遍 检验是否修改成功
}

func s() {
	//删：从数据表中删除一条数据
	var u1 UserInfo
	db.First(&u1) //从数据表查询第一条数据保存到u1
	//删除
	db.Delete(&u1)
}
