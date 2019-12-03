package loginRouters

import (
	"github.com/astaxie/beego"
	"mlt-go/module/login/controller"
)

func LoginRouter() {

	beego.Router("/login", &controller.LoginCotroller{}, "post:Login")

	beego.Router("/logout", &controller.LoginCotroller{}, "post:Logout")
	//添加用户
	beego.Router("/reg", &controller.RegController{}, "post:Reg")
}
