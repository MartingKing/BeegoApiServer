package models

import (
	_ "github.com/go-sql-driver/mysql" //初始化数据库
	"github.com/astaxie/beego/orm"
	"DHDApi/crawer"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type CityList struct {
	Id       int    `json:"id,omitemty"`
	Cityname string `json:"cityname,omitemty"`
	Cityurl  string `json:"cityurl,omitemty"`
}

const (
	aliasName  = "default"
	driverName = "mysql"
	dataSource = "root:huadong4305@tcp(127.0.0.1:3306)/test?charset=utf8"
)

func init() {
	orm.RegisterModel(new(CityList))
	ConnectSql()
}

//数据库插入city的信息
func AddCity() []CityList {
	var citylist []CityList
	matchesitem := crawer.Start()
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		panic(err)
	}
	index := 1
	for _, item := range matchesitem {
		//根据输出格式打印所需要的结果
		index++
		stmt, err := db.Prepare("INSERT citylist  SET id=?, cityname=?,cityurl=?")
		if err != nil {
			panic(err)
		}
		res, err := stmt.Exec(index, item[2], item[1])
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("insert success,city id:", id)
		}
	}
	orm.NewOrm().Raw("select * from citylist").QueryRows(&citylist)
	return citylist
}

func DeleteCity(uid string) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM citylist WHERE id=?", uid)
}

func GetCityById(id string) []CityList {
	var cl []CityList
	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		log.Fatal(err)
	}
	qb.Select("id", "cityname", "cityurl").From("citylist").Where("id=?").Limit(1)
	sq := qb.String()
	orm.NewOrm().Raw(sq, id).QueryRows(&cl)
	return cl
}
func GetAllCitys() []CityList {
	var s []CityList
	orm.NewOrm().Raw("select * from citylist").QueryRows(&s)
	return s
}

func GetPageCitys(pagecount string) []CityList {
	var cl []CityList
	var p, offset int
	qb, err := orm.NewQueryBuilder("mysql")
	p, err = strconv.Atoi(pagecount)
	if err != nil {
		log.Fatal(err)
	}
	//查询第一页 1-10条数据  第二页就是11-20
	offset = 10*p + 1
	qb.Select("id", "cityname", "cityurl").From("citylist").Limit(offset).Offset(10)
	sq := qb.String()
	orm.NewOrm().Raw(sq).QueryRows(&cl)
	return cl
}

func ConnectSql() {
	orm.RegisterDataBase(aliasName, driverName, dataSource, 30)
	orm.RunSyncdb(aliasName, false, true)
}
