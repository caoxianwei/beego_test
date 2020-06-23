package routers

import (
	"beego_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/user", &controllers.UserController{})
    beego.Router("/user/?:id", &controllers.UserController{}, "get:GetInfo")
}
