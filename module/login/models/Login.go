package models

type Login struct {
	//登录类型
	Type string
	Username string
	Password string
	Code string
}