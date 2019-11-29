package usersRouters

import (
	controller2 "mlt-go/module/users/controller"
	"github.com/astaxie/beego"
)

func UserRouter() {
	//添加用户
	beego.Router("/user/addUser", &controller2.UsersController{}, "post:AddUser")
}

