package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mlt-go/module/users/models"
	service2 "mlt-go/module/users/service"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) AddUser() {
	var users models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &users)
	c.Data["json"] = service2.AddUserService(users)
	//c.TplName = "index.tpl"
	c.ServeJSON()
}

func (c *UsersController) UserInfo() {

	token := c.GetString("token")
	c.Data["json"] = service2.UserInfo(token)
	c.ServeJSON()
}
