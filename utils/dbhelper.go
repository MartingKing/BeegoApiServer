package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //初始化数据库
)
const (
	AliasName  = "default"
	DriverName = "mysql"
	DataSource = "root:huadong4305@tcp(127.0.0.1:3306)/test?charset=utf8"
)

func ConnectSql() {
	orm.RegisterDataBase(AliasName, DriverName, DataSource, 30)
	orm.RunSyncdb(AliasName, false, true)
}
