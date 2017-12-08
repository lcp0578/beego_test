package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/lcp0578/beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lcp0578/beego_test/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUsername",
			Router: `/get_username`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

}
