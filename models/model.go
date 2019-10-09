package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "go:go@tcp(mianban.dycyec.com:3306)/go")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
