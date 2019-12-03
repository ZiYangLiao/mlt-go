package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mlt-go/common/models"
	"mlt-go/common/utils"
	models2 "mlt-go/module/login/models"
	models3 "mlt-go/module/users/models"
	"unsafe"
)

func Reg(reg models2.Reg) models.Result {
	//验证信息
	result := regValidate(reg)
	if result.Code == 0 {
		return result
	}
	reg.Password = utils.Md5Utils(reg.Password)
	reg.CreatDate = utils.GetTimeNow()

	//添加用户
	users := new(models3.Users)
	users.Email = reg.Email
	users.Mobile = reg.Mobile
	users.CreatDate = reg.CreatDate
	users.Realname = reg.Realname
	users.Nickname = reg.Nickname
	users.Birthday = reg.Birthday
	users.Age = reg.Age
	users.IDCard = reg.IDCard
	users.Status = 1
	o := orm.NewOrm()
	resultId, err := o.Insert(users)
	if err != nil {
		return models.Result{Code: 0, Msg: "注册失败"}
	}
	users.Id = resultId
	//添加登录鉴权帐号
	var account models2.Account
	account.Password = reg.Password
	account.CreatDate = reg.CreatDate
	account.Mobile = reg.Mobile
	account.Email = reg.Email
	account.UserId = resultId
	o = orm.NewOrm()
	_, err = o.Insert(&account)
	if err != nil {
		fmt.Println(err)
		o = orm.NewOrm()
		o.Delete(users)
		return models.Result{Code: 0, Msg: "注册失败"}
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
	utils.SetRedisEXKey(fmt.Sprintf(utils.Token+"user:%d", users.Id), key, int64(utils.ExpireMonth))
	utils.SetNXRedisEXKey(key, data, int64(utils.ExpireMonth))
	fmt.Println("redis.user.key:" + utils.Token + uuid)
	jsonBytes, err := json.Marshal(users)
	m := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &m)
	m["token"] = uuid
	return models.Result{Code: 1, Msg: "成功", Data: m}
}

func regValidate(reg models2.Reg) models.Result {

	if reg.Email == "" && reg.Mobile == "" {
		return models.Result{Code: 0, Msg: "邮箱或者手机号必须输入一个"}
	}
	if reg.Email != "" {
		if !utils.VerifyEmailFormat(reg.Email) {
			return models.Result{Code: 0, Msg: "邮箱格式不正确"}
		}

		o := orm.NewOrm()
		//err := o.QueryTable(users).Filter("email", reg.Email).One(&users)
		//if err != orm.ErrNoRows {
		//
		//}
		exist := o.QueryTable("users").Filter("email", reg.Email).Exist()
		if exist {
			return models.Result{Code: 0, Msg: "邮箱已存在"}
		}
	}
	if reg.Mobile != "" {
		if !utils.VerifyMobileFormat(reg.Mobile) {
			return models.Result{Code: 0, Msg: "手机号格式不正确"}
		}
		o := orm.NewOrm()
		exist := o.QueryTable("users").Filter("mobile", reg.Mobile).Exist()
		if exist {
			return models.Result{Code: 0, Msg: "手机号已存在"}
		}
	}
	if reg.Nickname == "" {
		return models.Result{Code: 0, Msg: "请输入用户名称"}
	}
	if len(reg.Nickname) > 20 {
		return models.Result{Code: 0, Msg: "用户名称长度不能超过20"}
	}
	if reg.Password == "" {
		return models.Result{Code: 0, Msg: "请输入密码"}
	}
	if len(reg.Password) < 6 || len(reg.Password) > 32 {
		return models.Result{Code: 0, Msg: "密码长度在6-32位"}
	}

	if reg.Realname != "" && len(reg.Realname) > 20 {
		return models.Result{Code: 0, Msg: "用户真实名称长度不能超过20"}
	}

	return models.Result{Code: 1, Msg: "成功"}
}
