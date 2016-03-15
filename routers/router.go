package routers

import (
	"github.com/shwinpiocess/cc/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "*:Index")
    beego.Router("/login", &controllers.MainController{}, "*:Login")
    beego.Router("/logout", &controllers.MainController{}, "*:Logout")
    
    // 业务管理
    beego.Router("/app/index", &controllers.AppController{}, "*:Index")
    beego.Router("/app/newapp", &controllers.AppController{}, "*:NewApp")
    beego.Router("/app/add", &controllers.AppController{}, "*:AddApp")
    beego.Router("/app/delete", &controllers.AppController{}, "*:DeleteApp")
}
