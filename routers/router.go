package routers

import (
	"beego_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/user", &controllers.UserController{})
    beego.Router("/user/?:id", &controllers.UserController{}, "get:GetInfo")
    beego.Router("/user/([0-9]?):id", &controllers.UserController{}, "get:GetNum;post:PostData")
}
