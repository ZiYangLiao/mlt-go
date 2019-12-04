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
    //状态  1 正常  2 冻结  3  删除
    Status int16
    //用户头像
    HeadPicUrl string
}

func init()  {
	orm.RegisterModel(new (Users))
}