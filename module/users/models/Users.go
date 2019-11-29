package models

import "github.com/astaxie/beego/orm"

type Users struct {
    Id int64
    Mobile string
    Email string
    Nickname string
    Realname string
    Age int16
	Birthday string
    IDCard string
	CreatDate string
}

func init()  {
	orm.RegisterModel(new (Users))
}