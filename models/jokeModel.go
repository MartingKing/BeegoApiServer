package models

import (
	"github.com/astaxie/beego/orm"
	"CrowerApi/utils"
	"database/sql"
	"fmt"
	"CrowerApi/crawer"
)

type Duanzi struct {
	Id       int    `json:"id"`
	Tittle   string `json:"tittle"`
	Contents string `json:"contents"`
}

func init() {
	orm.ResetModelCache()
	orm.RegisterModel(new(Duanzi))
	utils.ConnectSql()
}

func AddDuanzi() {
	crawer.DuanziSpiderStart(1, 10)
	db, err := sql.Open(utils.DriverName, utils.DataSource)
	if err != nil {
		panic(err)
	}
	index := 0
	// crawer.DuanziMap 保存的key是title  value是content
	for title, content := range crawer.DuanziMap {
		//根据输出格式打印所需要的结果
		index++
		stmt, err := db.Prepare("INSERT duanzilist  SET id=?, tittle=?,contents=?")
		if err != nil {
			panic(err)
		}
		res, err := stmt.Exec(index, title, content)
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("insert success,id:", id)
		}
	}
}
func DeleteAll() {
	orm.NewOrm().Raw("delete from duanzilist where 1=1")
}
func GetAllJoke() []Duanzi {
	var d []Duanzi
	orm.NewOrm().Raw("select * from duanzilist").QueryRows(&d)
	return d
}
