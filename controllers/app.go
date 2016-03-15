package controllers

import (
	"fmt"
	"strings"
    "github.com/shwinpiocess/cc/models"
)

type AppController struct {
    BaseController
}

func (this *AppController) Index() {
    num, apps, err := models.GetApps(this.userId)
    if err == nil {
        // 无任何业务时
        if num == 0 {
            this.TplName = "app/index_empty.html"
        } else {
            this.Data["apps"] = apps
            this.Data["defApp"] = apps[0]
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
        fmt.Println("uuuuuuuuuuuuuuuuuu")
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