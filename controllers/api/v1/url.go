package v1

import (
	"strconv"

	"github.com/TeamFat/FatUrl-Jsser/controllers/api"
	"github.com/TeamFat/FatUrl-Jsser/utils"
)

type UrlController struct {
	api.BaseController
}

func (url *UrlController) Index() {

}

func (this *UrlController) Encode() {
	paramId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	reponseData := make(map[string]interface{})
	reponseData["code"] = 200
	reponseData["data"] = map[string]interface{}{
		"id":     paramId,
		"encode": utils.HashIdEncode(paramId),
	}
	this.Data["json"] = &reponseData
	this.ServeJSON()
}

func (this *UrlController) Decode() {

	str := this.Ctx.Input.Param(":hash")
	reponseData := make(map[string]interface{})
	reponseData["code"] = 200
	reponseData["data"] = map[string]interface{}{
		"hash": str,
		"id":   utils.HashIdDecode(str),
	}
	this.Data["json"] = &reponseData
	this.ServeJSON()
}
