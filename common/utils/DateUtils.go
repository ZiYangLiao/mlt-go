package utils

import (
	"time"
)

func GetTimeNow() string {
	formatStr := "2006-01-02 15:04:05"
	return time.Now().Format(formatStr)
}
