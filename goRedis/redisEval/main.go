package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type MyRedisClient struct {
	client *redis.Client
}

const CONST_DISTRIBUTED_KEY = "d_lock"
const CONST_DISTRIBUTED_VALUE = "locked"
const CONST_DISTRIBUTED_TIMEOUT = 10

func (c *MyRedisClient) getDistributedLock(key string, value string, timeout int) int {
	luaScript := `
			local result = 0
			if redis.call("GET", KEYS[1]) == ARGV[1] then
				result = -1
			else
				local res = redis.call("SETEX", KEYS[1], ARGV[2], ARGV[1])
				for key,value in pairs(res) 
					do
						if key == "ok" then
							result = 1
						else
							result = -2
						end
				end
			end
			return result
		`
	result, err := c.client.Eval(luaScript, []string{key}, value, timeout).Int()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if result == 1 {
		return 1
	} else {
		return 0
	}
}
func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "139.198.9.163:6379", // Redis服务器地址和端口
		Password: "woshiwx@123",        // Redis密码，没有可以设置为空字符串
		DB:       0,                    // 使用的数据库，默认为0
	})
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("err connect")
		return
	}
	redisClient := &MyRedisClient{
		client: client,
	}
	res := redisClient.getDistributedLock(CONST_DISTRIBUTED_KEY, CONST_DISTRIBUTED_VALUE, CONST_DISTRIBUTED_TIMEOUT)
	if res == 1 {
		fmt.Println("locked")
	} else {
		fmt.Println("unlock")
	}
}
