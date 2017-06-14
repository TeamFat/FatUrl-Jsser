package controllers

import (
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
