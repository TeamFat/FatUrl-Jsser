package frontend

import (
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userId         int
	userName       string
}

//前置处理
func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	//登陆验证
	this.auth()

	this.Data["appName"] = beego.AppConfig.String("appname")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

//登陆验证
func (this *BaseController) auth() {

}

//渲染模板
func (this *BaseController) display(tpl ...string) {

}

//重定向
func (this *BaseController) redirect() {

}
