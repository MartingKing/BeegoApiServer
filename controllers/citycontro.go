package controllers

import (
	"CrowerApi/models"
	"github.com/astaxie/beego"
	"CrowerApi/utils"
)

// Operations about Citys
type CityController struct {
	beego.Controller
}

// @router /getallcity [get]
func (this *CityController) GetAll() {
	citys := models.GetAllCitys()
	this.Data["json"] = &citys
	this.ServeJSON()
}

// @router /getpagecity/:offset [get]
func (this *CityController) GetPageCity() {
	o := this.GetString(":offset")
	citys := models.GetPageCitys(o)
	this.Data["json"] = &citys
	this.ServeJSON()
}

// @router /getsinglecity/:cityid [get]
func (this *CityController) GetCity() {
	uid := this.GetString(":cityid")
	id := utils.GetCityid(uid)
	cites := models.GetCityById(string(id))
	this.Data["json"] = &cites
	this.ServeJSON()
}

func (this *CityController) ShowAPIVersion() {
	this.Ctx.WriteString("version 1.0")
}
