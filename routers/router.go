// @APIVersion 1.0.0
// @Title crower Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact dhd671108@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"CrowerApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/test",
		beego.NSNamespace("/city",
			beego.NSInclude(&controllers.CityController{}),
		),
		beego.NSNamespace("/duanzi",
			beego.NSInclude(&controllers.DuanziController{}),
		),
	)
	beego.AddNamespace(ns)
}
