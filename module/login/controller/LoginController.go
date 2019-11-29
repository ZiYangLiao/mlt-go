package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mlt-go/module/login/models"
	"mlt-go/module/login/service"
)

type LoginCotroller struct {
	beego.Controller
}

func (c *LoginCotroller) Login()  {
	var login models.Login
	json.Unmarshal(c.Ctx.Input.RequestBody, &login)
	c.Data["json"] = service.LoginService(login)
	c.ServeJSON()
}