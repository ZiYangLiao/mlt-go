package main

import (
	"github.com/astaxie/beego"
	"mlt-go/common/db"
	_ "mlt-go/routers"
)

func main() {
	//数据库连接初始化
	db.Init()
	beego.Run()
}

