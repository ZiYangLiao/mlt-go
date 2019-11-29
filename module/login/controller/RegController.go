package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mlt-go/module/login/models"
	"mlt-go/module/login/service"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) Reg()  {
	var reg models.Reg
	json.Unmarshal(c.Ctx.Input.RequestBody, &reg)

	c.Data["json"] = service.Reg(reg)
	c.ServeJSON()
}