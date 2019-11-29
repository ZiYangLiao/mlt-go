package utils

import (
	"crypto/md5"
	"fmt"
)

const md5Key  = "MLT-GO"

func Md5(data string) string  {
	return fmt.Sprintf("%x", md5.Sum([]byte(md5Key + data)))
}
