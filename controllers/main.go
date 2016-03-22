package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"

	"github.com/shwinpiocess/cc/models"
	"github.com/shwinpiocess/cc/utils"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	fmt.Println("this.firstApp=", this.firstApp)
	if this.firstApp {
		this.TplName = "app/help.html"
	} else {
		this.TplName = "index.html"
	}
}

// 登录
func (this *MainController) Login() {
	if this.userId > 0 {
		fmt.Println("登录userid=", this.userId)
		this.redirect("/")
	}

	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if username != "" && password != "" {
			user, err := models.GetUserByName(username)
			errMsg := ""
			if err != nil || user.Password != utils.Md5([]byte(password+user.Salt)) {
				errMsg = "账号或密码错误"
			} else if user.Status == -1 {
				errMsg = "该账号已禁用"
			} else {
				user.LastIp = this.getClientIP()
				user.LastLogin = time.Now().Unix()
				models.UpdateUser(user)

				authkey := utils.Md5([]byte(this.getClientIP() + "|" + user.Password + user.Salt))

				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
				}
				fmt.Println(this.Data)
				this.redirect(beego.URLFor("AppController.Index"))
			}
			fmt.Println(errMsg)
		}

	}

	fmt.Println("登录嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻嘻")
	this.TplName = "main/login.html"
}

// 退出登录
func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("MainController.Login"))
}
