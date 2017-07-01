package routers

import (
	"github.com/TeamFat/FatUrl-Jsser/controllers/api/v1"
	"github.com/TeamFat/FatUrl-Jsser/controllers/frontend"
	"github.com/astaxie/beego"
)

func init() {

	/* Api */

	beego.Router("/api/v1/url/encode", &v1.UrlController{}, "*:Encode")
	beego.Router("/api/v1/url/decode", &v1.UrlController{}, "*:Decode")

	/* Frontend */
	//首页
	beego.Router("/", &frontend.SiteController{}, "*:Index")
	//关于我们
	beego.Router("/about", &frontend.SiteController{}, "*:About")
	/* Backend */
}
