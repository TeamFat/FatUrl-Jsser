package frontend

/*
 *首页统一配置
 */
type SiteController struct {
	BaseController
}

func (this *SiteController) Prepare() {
	//重载父类方法
	this.BaseController.Prepare()
	this.Layout = "frontend/layout/layout.html"
}

//首页
func (this *SiteController) Index() {
	this.Data["IsHome"] = true
	this.TplName = "frontend/site/index.html"
}

//登陆
func (this *SiteController) Login() {
	this.Data["IsLogin"] = true

}

//注册
func (this *SiteController) Register() {
	this.Data["IsRegister"] = true

}

//关于我们
func (this *SiteController) About() {
	this.Data["IsAbout"] = true
	this.TplName = "frontend/site/about.html"
}
