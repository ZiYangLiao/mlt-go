package controllers

import (
	"fmt"
	models2 "mlt-go/common/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	GetBody := c.GetString("a")
	fmt.Println(GetBody)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	result := models2.Result{Code: 1, Msg: "成功"}
	c.Data["json"] = result
	//c.TplName = "index.tpl"
	c.ServeJSON()
}


