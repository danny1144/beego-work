package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"rest-api/models"
	_ "rest-api/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	u.Data["json"] = map[string]interface{}{"name": "astaxie"}
	u.ServeJSON()
}

// Post
func (u *UserController) Post() {
	log.Print("post参数", string(u.Ctx.Input.RequestBody))
	var ob models.User
	var err error
	if err = json.Unmarshal(u.Ctx.Input.RequestBody, &ob); err == nil {
		fmt.Print("参赛")
		o := orm.NewOrm()
		id, err := o.Insert(&ob)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("插入ID", id)

	}
	u.Data["json"] = "insert success"
	u.ServeJSON()

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
