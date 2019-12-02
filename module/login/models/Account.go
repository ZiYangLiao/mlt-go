package models

import "github.com/astaxie/beego/orm"

type Account struct {
	Id int64
	Password string
	Mobile string
	Email string
	Openid string
	Unionid int16
	UserId int64
	CreatDate string
}

func init()  {
	orm.RegisterModel(new (Account))
}