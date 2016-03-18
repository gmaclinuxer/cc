package controllers

import (
	"strings"
	"net/url"
	"strconv"
	"fmt"
	
	"github.com/shwinpiocess/cc/models"
)

type AppController struct {
	BaseController
}

func (this *AppController) Index() {
	fmt.Println("--------------------------------------------->")
	fmt.Println(this.Data)
	if this.appCount > 0 {
		this.TplName = "app/index.html"
	} else {
		this.TplName = "app/help.html"
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
		
		if Id, err := models.AddApp(app); err != nil {
			fmt.Println("err=", err)
			out["errInfo"] = "同名的业务已经存在！"
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} else {
			this.Ctx.SetCookie("defaultAppId", strconv.Itoa(Id))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(app.ApplicationName))
			
			var fields []string
			var sortby []string
			var order []string
			var query map[string]string = make(map[string]string)
			var limit int64 = 0
			var offset int64 = 0
		
			query["owner_id"] = strconv.Itoa(this.userId)
		
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
	this.TplName = "topology/set.html"
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

// 快速分配
func (this *AppController) QuickImport() {
	this.TplName = "host/quickImport.html"
}
