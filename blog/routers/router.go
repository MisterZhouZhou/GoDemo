package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/username", &controllers.UserController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/json", &controllers.JSONController{})

}
