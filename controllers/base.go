package controllers

import (
	"net/url"
	"fmt"
	"time"
	"strconv"
	"strings"
	
	"github.com/astaxie/beego"

	"github.com/shwinpiocess/cc/models"
	"github.com/shwinpiocess/cc/utils"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.User
	userId         int
	userName       string
	pageSize       int
	appCount       int
}

func (this *BaseController) Prepare() {
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()

	this.Data["path"] = this.Ctx.Request.RequestURI
	this.Data["today"] = time.Now().Format("20060102")
	
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	query["owner_id"] = strconv.Itoa(this.userId)
	fmt.Println("query=", query)

	apps, _ := models.GetAllApp(query, fields, sortby, order, offset, limit)
	this.appCount = len(apps)
	
	var defaultAppId int
	var defaultAppName string
	if this.appCount > 0 {
		defaultAppId, _ = strconv.Atoi(this.Ctx.GetCookie("defaultAppId"))
		defaultAppName =  this.Ctx.GetCookie("defaultAppName")
		if app, _:= models.GetAppById(defaultAppId); app == nil {
			defaultApp := apps[0].(models.App)
			defaultAppId = defaultApp.Id
			defaultAppName = defaultApp.ApplicationName
			this.Ctx.SetCookie("defaultAppId", strconv.Itoa(defaultAppId))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(defaultAppName))
		} else {
			defaultAppName, _ = url.QueryUnescape(defaultAppName)
		}
	}
	this.Data["defaultAppId"] = defaultAppId
	this.Data["defaultAppName"] = defaultAppName
	this.Data["apps"] = apps
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
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
				this.user = user
			}
		}
	}

	if this.userId == 0 && (this.controllerName != "main" || (this.controllerName == "main" && this.actionName != "logout" && this.actionName != "login")) {
		this.redirect(beego.URLFor("MainController.Login"))
	}
}
