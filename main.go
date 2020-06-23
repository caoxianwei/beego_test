package main

import (
	"beego_test/models"
	_ "beego_test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func inserUser()  {
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "ffforin"
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert err")
		return
	}
	beego.Info("insert success, id =", id)
}

func main() {
	inserUser()
	beego.Run()
}

