package controllers

import (
	"github.com/astaxie/beego"
	"github.com/oldlei/goblog/utils"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)Get(){
	this.TplName="login.tpl"
}
func (this *LoginController)Post(){
	name:=this.Input().Get("username")
	user := utils.User{}
	user.Name = name
	user.Id = "230"


	token := utils.GenerateToken(user,0)

	this.Ctx.WriteString(token)
}