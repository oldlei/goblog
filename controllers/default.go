package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/oldlei/goblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	err := o.Read(&user)
	if err != nil {
		fmt.Println(err)
	} else {
		logs.Info(user)
	}
	logs.Info(c.Ctx.Input.Param("name"))
	c.SetSession("uid", 20)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post() {
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "zhangsan"
	user.Password = "1111111"
	a, err := o.Insert(&user)
	if err != nil {
		logs.Info(err)
	} else {
		logs.Info("===")
		logs.Info(a)
	}
}
