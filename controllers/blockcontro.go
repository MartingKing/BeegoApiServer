package controllers

import (
	"CrowerApi/models"
	"CrowerApi/utils"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

type BlockController struct {
	beego.Controller
}

// @router /startInvest [post]
func (this *BlockController) Invest() {
	jsoninfo := this.GetString("amount")
	if jsoninfo == "" {
		this.Ctx.WriteString("amount is empty")
		return
	}
	this.insertDataToSw234(jsoninfo)
	this.Data["json"] = jsoninfo
	this.ServeJSON()
	fmt.Println("money==" + jsoninfo)
}

func (this *BlockController) insertDataToSw234(money string) {
	db, err := sql.Open(utils.DriverName, utils.DataSource)
	stmt, err := db.Prepare("INSERT block_miner  SET id=?, tittle=?,contents=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(2, "kim", "kim invested  "+money+" RMB")
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("insert success,miner id:", id)
	}
}

// @router /getInvestList [get]
func (this *BlockController) getInvestData() {
	invest := models.SearchInvestData()
	this.Data["json"] = &invest
	this.ServeJSON()
}
