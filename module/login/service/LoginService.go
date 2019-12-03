package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mlt-go/common/models"
	"mlt-go/common/utils"
	models2 "mlt-go/module/login/models"
	models3 "mlt-go/module/users/models"
	"mlt-go/module/users/service"
	"unsafe"
)

func LoginService(login models2.Login) models.Result {

	return validateLogin(login)
}

func LogoutService(token string) models.Result {
	redisUsers := service.GetUserModel(token)
	utils.DeleteRedisKey(fmt.Sprintf(utils.Token+"user:%d", redisUsers.Id))
	utils.DeleteRedisKey(utils.Token + token)
	return models.Result{Code: 1, Msg: "退出登录成功"}
}

func validateLogin(login models2.Login) models.Result {
	if login.Username == "" {
		return models.Result{Code: 0, Msg: "请输入用户名"}
	}
	if login.Password == "" {
		return models.Result{Code: 0, Msg: "请输入密码"}
	}
	if len(login.Password) < 6 || len(login.Password) > 32 {
		return models.Result{Code: 0, Msg: "密码长度在6-32位"}
	}
	if !utils.VerifyEmailFormat(login.Username) && !utils.VerifyMobileFormat(login.Username) {
		return models.Result{Code: 0, Msg: "用户名或密码不正确"}
	}
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("email", login.Username).Or("mobile", login.Username)

	var account models2.Account
	qs := o.QueryTable(account) //.Filter("email", login.Username).Filter("mobile", login.Username).One(&account)
	qs.SetCond(cond1).One(&account)
	//if err != orm.ErrNoRows {
	//	return models.Result{Code: 0, Msg: "用户名或密码不正确"}
	//}
	if account.Password != utils.Md5Utils(login.Password) {
		return models.Result{Code: 0, Msg: "用户名或密码不正确"}
	}
	o = orm.NewOrm()
	users := new(models3.Users)
	err := o.QueryTable("users").Filter("id", account.UserId).One(users)
	if err != nil {
		fmt.Println(err)
		return models.Result{Code: 0, Msg: "登录失败"}
	}

	Len := unsafe.Sizeof(*users)
	testBytes := &models.SliceMock{
		Addr: uintptr(unsafe.Pointer(users)),
		Cap:  int(Len),
		Len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))

	var uuid = utils.UUID()
	var key = (utils.Token + uuid)
	utils.DeleteRedisKey(utils.GetRedisKey(fmt.Sprintf(utils.Token+"user:%d", users.Id)))
	utils.SetRedisEXKey(fmt.Sprintf(utils.Token+"user:%d", account.UserId), key, int64(utils.ExpireMonth))
	utils.SetNXRedisEXKey(key, data, int64(utils.ExpireMonth))
	jsonBytes, err := json.Marshal(users)
	m := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &m)
	m["Token"] = uuid
	return models.Result{Code: 1, Msg: "登录成功", Data: m}
}
