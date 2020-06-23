package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string `orm:size(100)`
}
func init()  {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/class_beego?charset=utf8")

	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}
