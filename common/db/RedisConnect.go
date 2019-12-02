package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedisConn  redis.Conn
func init() {
	var err error
	RedisConn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("Redis连接成功")
	//defer RedisConn.Close()
}
