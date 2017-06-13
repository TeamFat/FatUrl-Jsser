package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Layout = "frontend/layout/layout.html"
	this.TplName = "frontend/index.html"
}
