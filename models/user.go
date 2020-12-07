package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int `orm:auto`
	Name        string
}

type Profile struct {
	Id          int
	Age         int16
}

type Post struct {
	Id    int
	Title string
}

type Tag struct {
	Id    int
	Name  string
}

func init() {
	sqlConn := beego.AppConfig.String("sqlconn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", sqlConn)
	if err != nil{
		fmt.Print("连接失败")
	}
	fmt.Print("连接数据库成功！")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
	orm.RunSyncdb("default", false, true)
}