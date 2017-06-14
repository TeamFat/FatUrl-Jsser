package frontend

import "github.com/TeamFat/FatUrl-Jsser/controllers"

type SiteController struct {
	controllers.BaseController
}

func (this *SiteController) Get() {

	this.Layout = "frontend/layout/layout.html"
	this.TplName = "frontend/index.html"
}

//登陆
func (this *SiteController) Login() {

}

//注册
func (this *SiteController) Register() {

}

//关于我们
func (this *SiteController) About() {

}
