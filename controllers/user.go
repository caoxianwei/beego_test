package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this*UserController) Get() {
	this.Ctx.WriteString("getInfo resposce seccsess")
}
func (this*UserController) GetInfo() {
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("getInfo data. id ="+id)
}
func (this*UserController) GetNum() {
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("getInfo data. id ="+id)
}
func (this*UserController) PostData() {
	//id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("this is post function")
}