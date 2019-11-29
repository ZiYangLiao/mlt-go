package controller

import (
	"encoding/json"
	"mlt-go/module/users/models"
	service2 "mlt-go/module/users/service"
	"github.com/astaxie/beego"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) AddUser()  {
	var users models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &users)
	c.Data["json"] = service2.AddUserService(users)
	//c.TplName = "index.tpl"
	c.ServeJSON()
}


