package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"mlt-go/common/db"
)

var Token = "token:"
var ExpireMonth = (60 * 60 * 24 * 30)

/**
存放字符串
*/
func SetRedisKey(key string, value string) {
	_, err := db.RedisConn.Do("SET", key, value)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

/**
存放json
*/
func SetNXRedisKey(key string, value []byte) {
	_, err := db.RedisConn.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func SetRedisEXKey(key string, value string, ex int64) {
	_, err := db.RedisConn.Do("SET", key, value, "EX", ex)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func SetNXRedisEXKey(key string, value []byte, ex int64) {
	_, err := db.RedisConn.Do("SET", key, value, "EX", ex)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func GetRedisKey(key string) string {
	value, err := redis.String(db.RedisConn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		//fmt.Printf("Get mykey: %v \n", value)
	}
	return value
}

func GetNXRedisKey(key string) []byte {
	value, err := redis.Bytes(db.RedisConn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		//fmt.Printf("Get mykey: %v \n", value)
	}
	return value
}

func LPushRedisKey(key string, value string) {
	_, err := db.RedisConn.Do("lpush", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func LRangeRedisKey(key string, start int32, end int32) []interface{} {
	values, _ := redis.Values(db.RedisConn.Do("lrange", key, start, end))
	return values
}

func ExistsRedisKey(key string) bool {
	is_key_exit, err := redis.Bool(db.RedisConn.Do("EXISTS", key))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
	return is_key_exit
}

func DeleteRedisKey(key string) bool {
	_, err := db.RedisConn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delelte failed:", err)
		return false
	}
	return true
}
