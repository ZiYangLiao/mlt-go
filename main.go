package main

import (
	"mlt-go/common/db"
	_ "mlt-go/routers"
	"github.com/astaxie/beego"
)

func main() {
	//数据库连接初始化
	db.Init()

	beego.Run()
}

