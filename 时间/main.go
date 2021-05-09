package main

import (
	"fmt"
	"time"
)

/*
timer原理:通过无缓冲chan阻塞队列来实现延迟时间,即chan阻塞至timer初始化时所定义的时间
*/
func demoForTimer() {
	timer := time.NewTimer(time.Second * 3)
	<-timer.C
	fmt.Println("Hello,world")
	timer.Reset(time.Second * 2)
	<-timer.C
	fmt.Println("Hello,world")
}
func demoForTimerUseful() {
	var value chan int = make(chan int)
	timer := time.NewTimer(time.Second)
	go func() {
		time.Sleep(time.Second * 5)
		value <- 100
		close(value)
	}()
	select {
	case res := <-value:
		fmt.Printf("THE VALUE IS %d\n", res)
	case <-timer.C:
		fmt.Println("TIMEOUT")
	}
}
func demoForTicker() {
	ticker := time.NewTicker(time.Second)
	count := 0
	resetCount := 0
	for {
		<-ticker.C
		fmt.Println("HELLO,WORLD")
		count++
		if count == 3 {
			ticker.Reset(time.Second * 2)
			resetCount++
		}
		if resetCount == 2 {
			break
		}
	}
}

/*
It returns true if the call stops the timer, false if the timer has already
expired or been stopped.
如果timer已经停止或者过期,则认为此时timer.C中是存在值的,因此可以获取timer.C中的值
如果timer未过期,则会停止计时,也就不会再往timer.C中发送值
*/
func demoForStop() {
	timer := time.NewTimer(time.Second * 10)
	stopFlag := timer.Stop()
	fmt.Println(stopFlag)
	if !stopFlag {
		res := <-timer.C
		fmt.Println(res)
	} else {
		return
	}
}

/*
使用ticker控制for循环的单次循环时间
*/
func demoForLoopTimeControl() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		time.Sleep(20)
		fmt.Println(time.Now())
	}
}

func demoForTimeAfterFunctionCall() {
	sig := make(chan int)
	a := 1
	time.AfterFunc(time.Second, func() {
		a += 1
		sig <- 1
	})
	<-sig
	fmt.Println(a)
}

func review() {
	// 创建一个打点器, 每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)
	// 声明计数变量
	var i int
	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case temp := <-ticker.C: // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i, temp)
		case temp := <-stopper.C: // 计时器到时了
			fmt.Println("stop", temp)
			// 跳出循环
			goto StopHere

		}
	}
	// 退出的标签, 使用goto跳转
StopHere:
	fmt.Println("done")
}
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	timestampDemo2(timestamp2)
}
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func timeCalculate() {
	now := time.Now()
	later := now.Add(time.Hour)
	duration := time.Now().Sub(later)
	before := time.Now().Before(now.Add(time.Hour * 24 * 356))
	after := time.Now().After(now.Add(time.Hour * 24 * 356))

	fmt.Printf("%v\n", duration)
	fmt.Printf("%v\n", before)
	fmt.Printf("%v\n", after)
}
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05.000"))
}
func main() {
	demoForTimeAfterFunctionCall()
}
