package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegomyblog/controllers:MainController"] = append(beego.GlobalControllerRouter["beegomyblog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegomyblog/controllers:MainController"] = append(beego.GlobalControllerRouter["beegomyblog/controllers:MainController"],
		beego.ControllerComments{
			Method: "AbountMe",
			Router: `/aboutme`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegomyblog/controllers:MainController"] = append(beego.GlobalControllerRouter["beegomyblog/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetPost",
			Router: `/post/:file_name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
