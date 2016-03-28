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
	beego.Router("/app/getMainterners", &controllers.AppController{}, "*:GetMainterners")

	beego.Router("/topology/index", &controllers.AppController{}, "*:TopologyIndex")
	
	beego.Router("/Set/getAllSetInfo", &controllers.SetController{}, "*:GetAllSetInfo")
	beego.Router("/Set/newSet", &controllers.SetController{}, "*:NewSet")
	beego.Router("/Set/editSet", &controllers.SetController{}, "*:EditSet")
	beego.Router("/Set/delSet", &controllers.SetController{}, "*:DelSet")
	beego.Router("/Set/getSetInfoById", &controllers.SetController{}, "*:GetSetInfoById")
	
	beego.Router("/Module/newModule", &controllers.ModuleController{}, "*:NewModule")
	
	// 快速分配
	beego.Router("/host/quickImport", &controllers.HostController{}, "*:QuickImport")
	beego.Router("host/hostQuery", &controllers.HostController{}, "*:HostQuery")
	
	beego.Router("/host/getHostById", &controllers.HostController{}, "post:GetHostById")
	beego.Router("/host/importPrivateHostByExcel", &controllers.HostController{}, "post:ImportPrivateHostByExcel")
	beego.Router("/host/getHost4QuickImport", &controllers.HostController{}, "post:GetHost4QuickImport")
	beego.Router("/host/delPrivateDefaultApplicationHost", &controllers.HostController{}, "post:DelPrivateDefaultApplicationHost")
	beego.Router("/host/quickDistribute", &controllers.HostController{}, "post:QuickDistribute")
	beego.Router("/host/details", &controllers.HostController{}, "post:Details")
	beego.Router("/host/resHostModule/", &controllers.HostController{}, "post:ResHostModule")
	beego.Router("/host/getTopoTree4view", &controllers.HostController{}, "post:GetTopoTree4view")
	beego.Router("/host/modHostModule/", &controllers.HostController{}, "post:ModHostModule")
	beego.Router("/host/delHostModule/", &controllers.HostController{}, "post:DelHostModule")
}
