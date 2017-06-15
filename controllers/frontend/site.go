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

func (this *SiteController) Index() {

	this.TplName = "frontend/site/index.html"
}

//登陆
func (this *SiteController) Login() {

}

//注册
func (this *SiteController) Register() {

}

//关于我们
func (this *SiteController) About() {
	this.TplName = "frontend/site/about.html"
}
