package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"regexp"
)

const md5Key = "MLT-GO"

func Md5Utils(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(md5Key+data)))
}

func UUID() string  {
	return uuid.Must(uuid.NewV4()).String()
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyMobileFormat(mobile string) bool {
	pattern := `^(1[3|4|5|6|7|8|9][0-9]\d{4,8})$` //匹配手机号
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}
