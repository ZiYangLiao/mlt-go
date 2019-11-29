package service

import (
	"mlt-go/common/models"
	models2 "mlt-go/module/login/models"
)

func Reg(reg models2.Reg) models.Result  {

	return models.Result{Code: 1, Msg: "成功"}
}
