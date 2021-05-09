package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func initDB() *redis.Client {
	option := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	db := redis.NewClient(option)
	return db
}
func checkConnectionOfDB(db *redis.Client) bool {
	_, err := db.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func closeDB(db *redis.Client) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

/*
SetNX()与SetEX()的区别是，SexNX()仅当key不存在的时候才设置，
如果key已经存在则不做任何操作，而SetEX()方法不管该key是否已经存在缓存中直接覆盖
*/
/*
我们想知道我们调用SetNX()是否设置成功了，可以接着调用Result()方法，返回的第一个值表示是否设置成功了，如果返回false,说明缓存Key已经存在，
此次操作虽然没有错误，但是是没有起任何效果的。如果返回true，表示在此之前key是不存在缓存中的，操作是成功的
*/
/*
如果要获取的key在缓存中并不存在，Get()方法将会返回redis.Nil
*/
/*
Incr()、IncrBy()都是操作数字，对数字进行增加的操作，incr是执行原子加1操作，incrBy是增加指定的数
Decr()和DecrBy()方法是对数字进行减的操作，和Incr正好相反
*/
func stringSetOperationForDB(db *redis.Client) {
	opr := db.Set("name", "Hydra2", 0)
	res, err := opr.Result()
	fmt.Println(res, err)
	opr2 := db.SetNX("name", "HydraCode", 0)
	validity, err := opr2.Result()
	fmt.Println(validity, err)
}
func getOperationForDB(db *redis.Client) {
	getValue := db.Get("name")
	if getValue == nil {
		fmt.Println("The value of the key 'name' is nil")
	} else {
		fmt.Printf("The value of the key 'name' is %s\n", getValue)
	}
}
func doFunctionForDB(db *redis.Client) {
	cmd := db.Do(ctx, "set1", "name", "Hydra")
	res, err := cmd.Result()
	fmt.Println(res, err)
}

func main() {

}
