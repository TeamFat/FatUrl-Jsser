package routers

import (
	"github.com/TeamFat/FatUrl-Jsser/controllers/frontend"
	"github.com/astaxie/beego"
)

func init() {
	/* Frontend */
	//首页
	beego.Router("/", &frontend.SiteController{}, "*:Index")
	//关于我们
	beego.Router("/about", &frontend.SiteController{}, "*:About")
	/* Backend */
}
