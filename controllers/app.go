package controllers

import (
	"net/url"
	"strconv"
	"fmt"
	"github.com/shwinpiocess/cc/models"
	"strings"
	"time"
)

type AppController struct {
	BaseController
}

func (this *AppController) Index() {
	this.Data["today"] = time.Now().Format("20060102")
	num, apps, err := models.GetApps(this.userId)
	if err == nil {
		// 无任何业务时
		if num == 0 {
			this.TplName = "app/index_empty.html"
		} else {
			this.Data["apps"] = apps
			if defaultAppId, err := strconv.Atoi(this.Ctx.GetCookie("defaultAppId")); err == nil {
				this.Data["defaultAppId"] = defaultAppId
				defaultAppName, err := url.QueryUnescape(this.Ctx.GetCookie("defaultAppName"))
				fmt.Println("err--->", err)
				fmt.Println("000000000000xxxxxxx", this.Ctx.GetCookie("defaultAppName"))
				this.Data["defaultAppName"] = defaultAppName
				fmt.Println("777777777777777777777777777777777777", this.Data["defaultAppId"], defaultAppName)
			} else {
				fmt.Println("8888888888888888888888888888888")
				this.Data["defaultAppId"] = apps[0].Id
				this.Data["defaultAppName"] = apps[0].ApplicationName
			}
			this.TplName = "app/index.html"
		}
	}
}

func (this *AppController) NewApp() {
	this.TplName = "app/newapp.html"
}

func (this *AppController) AddApp() {
	if this.isPost() {
		app := new(models.App)
		app.Type, _ = this.GetInt("Type")
		app.Level, _ = this.GetInt("Level")
		app.ApplicationName = strings.TrimSpace(this.GetString("ApplicationName"))
		app.LifeCycle = strings.TrimSpace(this.GetString("LifeCycle"))
		app.OwnerId = this.userId
		out := make(map[string]interface{})
		if models.IsAppExistByName(app.ApplicationName) {
			out["errInfo"] = "同名的业务已经存在！"
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		}
		if _, err := models.AddApp(app); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0007"
			this.jsonResult(out)
		}
		
		this.Ctx.SetCookie("defaultAppId", strconv.Itoa(app.Id))
		this.Ctx.SetCookie("defaultAppName", url.QueryEscape(app.ApplicationName))
		
		cnt, err := models.GetAppCountByUserId(this.userId)
		fmt.Println("cnt=", cnt, "err=", err)
		if err == nil {
			if cnt > 1 {
				out["success"] = true
				out["gotopo"] = 0
				this.jsonResult(out)
			}

			out["success"] = true
			out["gotopo"] = 1
			this.jsonResult(out)
		}
		out["success"] = false
		out["errInfo"] = err.Error()
		out["errCode"] = "0008"
		this.jsonResult(out)
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
	this.TplName = "topology/index.html"
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
