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

func QueryUser()  {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		beego.Info("query is err")
		return
	}
	beego.Info("query success, user =", user)
}

func userUpdate()  {
	o := orm.NewOrm()
	user := models.User{Id:1, Name:"hahwwww"}
	_, err := o.Update(&user)
	if err != nil {
		beego.Info("update err")
		return
	}
	beego.Info("update success")
}

func userDelete()  {
	o := orm.NewOrm()
	user := models.User{Id:2}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("Delete err")
		return
	}
	beego.Info("Delete success")
}
func main() {
	inserUser()
	QueryUser()
	//userUpdate()
	//userDelete()
	beego.Run()
}

