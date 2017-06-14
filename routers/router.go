package routers

import (
	"github.com/TeamFat/FatUrl-Jsser/controllers/frontend"
	"github.com/astaxie/beego"
)

func init() {
	//frontend
	beego.Router("/", &frontend.SiteController{})

	//backend
}
