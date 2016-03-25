package controllers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/shwinpiocess/cc/models"
)

type AppController struct {
	BaseController
}

func (this *AppController) Index() {
	fmt.Println("--------------------------------------------->")
	fmt.Println(this.Data)
	if this.firstApp {
		this.TplName = "app/help.html"
	} else {
		this.TplName = "app/index.html"
	}
}

func (this *AppController) NewApp() {
	this.TplName = "app/newapp.html"
}

func (this *AppController) AddApp() {
	if this.isPost() {
		app := new(models.App)
		app.Type, _ = this.GetInt8("Type")
		app.Level, _ = this.GetInt8("Level")
		app.ApplicationName = strings.TrimSpace(this.GetString("ApplicationName"))
		app.LifeCycle = strings.TrimSpace(this.GetString("LifeCycle"))
		app.OwnerId = this.userId

		out := make(map[string]interface{})

		if models.ExistByName(app.ApplicationName) {
			out["errInfo"] = "同名的业务已经存在！"
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		}
		
		if Id, err := models.AddApp(app); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} else {
			s := new(models.Set)
			s.ApplicationID = int(Id)
			s.SetName = "空闲机池"
			s.EnviType = 3
			s.ServiceStatus = 1
			s.Default = true
			s.Owner = this.userId
			if setId, err := models.AddSet(s); err == nil {
				m := new(models.Module)
				m.ApplicationId = int(Id)
				m.SetId = int(setId)
				m.ModuleName = "空闲机"
				m.Owner = this.userId
				models.AddModule(m)
			}
			
			models.AddDefApp(this.userId)
			
			this.Ctx.SetCookie("defaultAppId", strconv.FormatInt(Id, 10))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(app.ApplicationName))

			var fields []string
			var sortby []string
			var order []string
			var query map[string]string = make(map[string]string)
			var limit int64 = 0
			var offset int64 = 0

			query["owner_id"] = strconv.Itoa(this.userId)
			query["default"] = strconv.FormatBool(false)

			apps, _ := models.GetAllApp(query, fields, sortby, order, offset, limit)
			if len(apps) > 1 {
				out["success"] = true
				out["gotopo"] = 0
				this.jsonResult(out)
			} else {
				out["success"] = true
				out["gotopo"] = 1
				this.jsonResult(out)
			}
		}

		//		this.Ctx.SetCookie("defaultAppId", strconv.Itoa(app.Id))
		//		this.Ctx.SetCookie("defaultAppName", url.QueryEscape(app.ApplicationName))

		//		cnt, err := models.GetAppCountByUserId(this.userId)
		//		fmt.Println("cnt=", cnt, "err=", err)
		//		if err == nil {
		//			if cnt > 1 {
		//				out["success"] = true
		//				out["gotopo"] = 0
		//				this.jsonResult(out)
		//			}

		//			out["success"] = true
		//			out["gotopo"] = 1
		//			this.jsonResult(out)
		//		}
		//		out["success"] = false
		//		out["errInfo"] = err.Error()
		//		out["errCode"] = "0008"
		//		this.jsonResult(out)
	}
}

func (this *AppController) DeleteApp() {
	out := make(map[string]interface{})
	applicationId, _ := this.GetInt("ApplicationID")
	if err := models.DeleteApp(applicationId); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = err.Error()
		this.jsonResult(out)
	}
}

func (this *AppController) TopologyIndex() {
	
		topo, err := models.GetAppTopoById(this.defaultApp.Id)
		this.Data["topo"] = topo
		fmt.Println("topo=", topo, "err=", err)
	if this.defaultApp.Level == 3 {
		this.Data["desetid"] = 0
		this.TplName = "topology/set.html"
	} else {
		s := models.GetDesetidByAppId(this.defaultApp.Id)
		this.Data["desetid"] = s.SetID
		this.TplName = "topology/index.html"
	}
}

// 切换默认业务
func (this *AppController) SetDefaultApp() {
	out := make(map[string]interface{})
	if applicationId, err := this.GetInt("ApplicationID"); err != nil {
		out["success"] = false
		this.jsonResult(out)
	} else {
		if app, err := models.GetAppById(applicationId); err != nil {
			out["success"] = false
			this.jsonResult(out)
		} else {
			this.Ctx.SetCookie("defaultAppId", strconv.Itoa(applicationId))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(app.ApplicationName))
			fmt.Println("-------->", app.ApplicationName)
			out["success"] = true
			out["message"] = "业务切换成功"
			this.jsonResult(out)
		}
	}
}

func (this *AppController) GetMainterners() {
	var uinList = [...]int{10001, 10002}
	var UserNameList = [...]string{"张三", "李四"}
	var mainterners = make(map[string]interface{})
	mainterners["uinList"] = uinList
	mainterners["UserNameList"] = UserNameList

	this.Data["json"] = mainterners
	this.ServeJSON()
}