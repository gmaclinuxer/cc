package routers

import (
	"github.com/astaxie/beego"
	"github.com/shwinpiocess/cc/controllers"
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
	beego.Router("/welcome/setDefaultApp", &controllers.AppController{}, "*:SetDefaultApp")

	beego.Router("/topology/index", &controllers.AppController{}, "*:TopologyIndex")
}
