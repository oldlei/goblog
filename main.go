package main

import (
	"github.com/astaxie/beego"
	_ "goblog/models"
	_ "goblog/routers"
)

func main() {
	beego.Run()
}
