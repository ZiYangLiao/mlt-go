package routers

import (
	"mlt-go/controllers"
	loginRouters "mlt-go/module/login/loginRouters"
	"mlt-go/module/users/usersRouters"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	usersRouters.UserRouter();
	loginRouters.LoginRouter()
}
