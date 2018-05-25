package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["CrowerApi/controllers:CityController"] = append(beego.GlobalControllerRouter["CrowerApi/controllers:CityController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getallcity`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CrowerApi/controllers:CityController"] = append(beego.GlobalControllerRouter["CrowerApi/controllers:CityController"],
		beego.ControllerComments{
			Method: "GetPageCity",
			Router: `/getpagecity/:offset`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CrowerApi/controllers:CityController"] = append(beego.GlobalControllerRouter["CrowerApi/controllers:CityController"],
		beego.ControllerComments{
			Method: "GetCity",
			Router: `/getsinglecity/:cityid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
