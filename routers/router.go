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
	
	beego.Router("/App/getMainterners", &controllers.AppController{}, "*:GetMainterners")

	beego.Router("/topology/index", &controllers.AppController{}, "*:TopologyIndex")
	
	beego.Router("/Set/getAllSetInfo", &controllers.SetController{}, "*:GetAllSetInfo")
	beego.Router("/Set/newSet", &controllers.SetController{}, "*:NewSet")
	beego.Router("/Set/editSet", &controllers.SetController{}, "*:EditSet")
	beego.Router("/Set/delSet", &controllers.SetController{}, "*:DelSet")
	beego.Router("/Set/getSetInfoById", &controllers.SetController{}, "*:GetSetInfoById")
	
	beego.Router("/Module/newModule", &controllers.ModuleController{}, "*:NewModule")
	
	// 快速分配
	beego.Router("/host/quickImport", &controllers.AppController{}, "*:QuickImport")
}
