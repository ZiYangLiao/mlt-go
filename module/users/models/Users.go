package models

import "github.com/astaxie/beego/orm"

type Users struct {
    Id int64
    Name string
    Age int16
	Birthday string
	CreatDate string
}

func init()  {
	orm.RegisterModel(new (Users))
}