package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mlt-go/common/models"
	"mlt-go/common/utils"
	. "mlt-go/module/users/models"
	"unsafe"
)

func AddUserService(user Users) models.Result {

	user.CreatDate = utils.GetTimeNow()
	o := orm.NewOrm()
	addCount, _ := o.Insert(&user)
	fmt.Println(addCount)
	result := models.Result{Code: 0, Msg: "添加失败"}
	if addCount > 0 {
		result = models.Result{Code: 1, Msg: "成功"}
	}
	return result
}

func UserInfo(token string) models.Result {
	users := GetUserModel(token)
	if users == nil {
		return models.Result{Code: 0, Msg: "获取用户信息失败，请重新登录"}
	}
	jsonBytes, _ := json.Marshal(users)
	m := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &m)
	m["Token"] = token
	return models.Result{Code: 1, Msg: "操作成功", Data: m}
}

func GetUserModel(token string) *Users {
	userBytes := utils.GetNXRedisKey(utils.Token + token)
	var ptestStruct *Users = *(**Users)(unsafe.Pointer(&userBytes))
	return ptestStruct
}
