package routers

import (
	"github.com/TeamFat/FatUrl-Jsser/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
