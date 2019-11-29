package routers

import (
	"mlt-go/controllers"
	"mlt-go/module/users/usersRouters"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	usersRouters.UserRouter();
}
