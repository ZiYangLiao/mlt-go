package service

import (
	"fmt"
	"mlt-go/common/models"
	"mlt-go/common/utils"
	models3 "mlt-go/module/users/models"
	"github.com/astaxie/beego/orm"
)

func AddUserService(user models3.Users) models.Result  {

	user.CreatDate  = utils.GetTimeNow()
	o := orm.NewOrm();
	addCount, _ := o.Insert(&user)
	fmt.Println(addCount)
	result := models.Result{Code: 0, Msg: "添加失败"}
	if addCount > 0 {
		result = models.Result{Code: 1, Msg: "成功"}
	}
	return result
}
