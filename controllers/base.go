package controllers

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/shwinpiocess/cc/models"
	"github.com/shwinpiocess/cc/utils"
)

type BaseController struct {
	beego.Controller

	userId         int
	userName       string
	
	requestPath string
	today string
	
	firstApp bool
	firstSet bool
	firtModule bool
	
	defaultApp *models.App
}

func (this *BaseController) Prepare() {
	this.requestPath = this.Ctx.Input.URL()
	this.today = time.Now().Format("20060102")
	
	this.auth()

	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	query["owner_id"] = strconv.Itoa(this.userId)
	query["default"] = strconv.FormatBool(false)

	apps, _ := models.GetAllApp(query, fields, sortby, order, offset, limit)
	
	if len(apps) > 0 {
		defaultAppId := this.Ctx.GetCookie("defaultAppId")
		
		if defaultAppId, err := strconv.Atoi(defaultAppId); err == nil {
			if app, err := models.GetAppById(defaultAppId); err != nil {
				defaultApp := apps[0].(models.App)
				this.defaultApp = &defaultApp
				this.Ctx.SetCookie("defaultAppId", strconv.Itoa(this.defaultApp.Id))
				this.Ctx.SetCookie("defaultAppName", url.QueryEscape(this.defaultApp.ApplicationName))
			} else {
				this.defaultApp = app
			}
		} else {
			defaultApp := apps[0].(models.App)
			this.defaultApp = &defaultApp
			this.Ctx.SetCookie("defaultAppId", strconv.Itoa(this.defaultApp.Id))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(this.defaultApp.ApplicationName))
		}
	} else {
		this.firstApp = true
		this.firstSet = true
		this.firtModule = true
	}

	this.Data["requestPath"] = this.requestPath
	this.Data["today"] = this.today

	this.Data["defaultApp"] = this.defaultApp
	this.Data["apps"] = apps
	this.Data["firstApp"] = this.firstApp
	this.Data["firstSet"] = this.firstSet
	this.Data["firstModule"] = this.firtModule
}

func (this *BaseController) getClientIP() string {
	return strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) auth() {
	arrs := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arrs) == 2 {
		idstr, password := arrs[0], arrs[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.GetUserById(userId)
			if err == nil && password == utils.Md5([]byte(this.getClientIP()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id
				this.userName = user.UserName
			}
		}
	}

	if this.userId == 0 && this.requestPath != beego.URLFor("MainController.Login") && this.requestPath != beego.URLFor("MainController.Logout") {
		this.redirect(beego.URLFor("MainController.Login"))
	}
}
