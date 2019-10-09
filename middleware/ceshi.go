package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/oldlei/goblog/utils"
)
var FilterUser = func(ctx *context.Context) {
	a, ok := ctx.Input.Session("uid").(int)
	if !ok {
		ctx.Redirect(302, "/login")
		return
	}
	logs.Info(a)
	logs.Info("======111=======")
}
var FilterToken = func(ctx *context.Context) {
	logs.Info("current router path is ", ctx.Request.RequestURI)
	if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") == "" {
		logs.Error("without token, unauthorized !!")
		ctx.ResponseWriter.WriteHeader(401)
		ctx.ResponseWriter.Write([]byte("no permission"))
	}
	if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") != "" {
		token := ctx.Input.Header("Authorization")
		//logs.Info("-"+token+"-")
		//token = strings.Split(token, "")[1]
		logs.Info("curernttoken: ", token)
		logs.Info("=============================")
		error := utils.ValidateToken(token)
		if error != nil{
			logs.Info("出错了，",error)
		}else{
			logs.Info("验证通过")
		}
		logs.Info("=============================")
		// validate token
		// invoke ValidateToken in utils/token
		// invalid or expired todo res 401
	}
}

func init(){

	beego.InsertFilter("/*",beego.BeforeRouter,FilterToken)
}
