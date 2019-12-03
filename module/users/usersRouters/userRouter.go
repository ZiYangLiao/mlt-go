package usersRouters

import (
	"github.com/astaxie/beego"
	controller2 "mlt-go/module/users/controller"
)

func UserRouter() {
	//添加用户
	beego.Router("/user/addUser", &controller2.UsersController{}, "post:AddUser")

	beego.Router("/user/info", &controller2.UsersController{}, "Get:UserInfo")
}
