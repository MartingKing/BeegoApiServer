package controllers

import (
	"github.com/astaxie/beego"
	"CrowerApi/models"
)

type DuanziController struct {
	beego.Controller
}

// @router /getallduanzi [get]
func (this *DuanziController) GetAllDuanzi() {
	duane := models.GetAllJoke()
	this.Data["json"] = &duane
	this.ServeJSON()
}
