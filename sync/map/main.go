package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

/*
从map中取一个不存在的key值不会报错,可以通过第二个返回值validate验证返回值的有效性
从map中删除一个不存在的key不会报错
map本身是无序的，在遍历的时候并不会按照你传入的顺序，进行传出。
map声明之后必须初始化，才能使用
*/
func demoForRegularMap() {
	var myMap map[int]string
	myMap = make(map[int]string)
	res, validate := myMap[1]
	fmt.Println(res, validate)
	delete(myMap, 22)
	myMap[1] = "first"
	myMap[2] = "second"
	myMap[3] = "third"
	myMap[4] = "fourth"
	for key, value := range myMap {
		fmt.Printf("key is [%d] , value is [%s]\n", key, value)
	}
	var sortedKey []int
	for key := range myMap {
		sortedKey = append(sortedKey, key)
	}
	sort.Ints(sortedKey)
	for _, key := range sortedKey {
		fmt.Printf("key is [%d] , value is [%s]\n", key, myMap[key])
	}
}

/*
sync.Map的原理介绍：
sync.Map里头有两个map一个是专门用于读的read map，另一个是才是提供读写的dirty map；
优先读read map，若不存在则加锁穿透读dirty map，同时记录一个未从read map读到的计数，
当计数到达一定值，就将read map用dirty map进行覆盖。
sync.Map.Range()所接受的参数一个函数,该函数入参为两个类型为interface{}的变量,表示sync.Map的key和value
该函数的返回值是一个bool类型,如果返回false则停止后续的遍历
*/
func demoForSyncMap() {
	var mySyncMap = sync.Map{}
	mySyncMap.Delete(23)
	value, validate := mySyncMap.Load(23)
	fmt.Println(value, validate)
	mySyncMap.Store(1, "first")
	mySyncMap.Store("second", 2)
	mySyncMap.Range(
		func(key, value interface{}) bool {
			fmt.Printf("The key is %v,The value is %v\n", key, value)
			return true
		},
	)
}
func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
func demoForMapDeleteMechanism() {
	printMemStats("初始化")
	// 添加1w个map值
	intMap := make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		intMap[i] = i
	}
	// 手动进行gc操作
	runtime.GC()
	// 再次查看数据
	printMemStats("增加map数据后")
	fmt.Println("删除前数组长度：", len(intMap))
	for i := 0; i < 10000; i++ {
		delete(intMap, i)
	}
	fmt.Println("删除后数组长度：", len(intMap))
	// 再次进行手动GC回收
	runtime.GC()
	printMemStats("删除map数据后")

	// 设置为nil进行回收
	intMap = nil
	runtime.GC()
	printMemStats("设置为nil后")
}
func main() {
	demoForMapDeleteMechanism()
}
