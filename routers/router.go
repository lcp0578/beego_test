package routers

import (
	"github.com/lcp0578/beego_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.UserController{})
}
