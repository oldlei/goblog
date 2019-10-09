package routers

import (
	"github.com/astaxie/beego"
	"github.com/oldlei/goblog/controllers"
	_ "github.com/oldlei/goblog/middleware"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:id([0-9]+)", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
