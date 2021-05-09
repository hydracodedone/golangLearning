# channel
## deadlock 所有的groutine都等待则发生死锁
## 阻塞
装满以后如果继续输入则会阻塞该输入所在的groutine，并将cpu使用权交出，sleep
为空如果继续输出则会发生阻塞该输出所在的groutine，并将cpu使用权交出，sleep
##	并发
##	channel 
###	channel 数据接收的两种方式
    info <-channel 
    info,status <-channel
###	两种方式的区别在于
#### 第一种
如果channel没有关闭且channel中没有数据，则会阻塞该语句所在的channel
如果channel关闭，如果有数据，则返回数据，没有数据，则返回0
#### 第二种
如果channel关闭后，如果channel中的数据被读取完毕后，会返回 0，false
关闭channel后，如果channel中的数据没有被读取完，会返回数据，ture
没有关闭，有数据，会返回数据，true
没有关闭，没有数据，会阻塞
```
func demo2() {
	var myChan chan int
	myChan = make(chan int, 6)
	for i := 0; i < 6; i++ {
		myChan <- i
	}
	myChan <- 90
	fmt.Println("SS")
}
func main(){
    demo2()
}
```
分析：
    在main groutine 中，由于myChan已经装满，因此，myChan <- 90 会阻塞
    而整个程序只有一个groutine，因此当所有的groutine都sleep，因此DEAD LOAC 程序崩溃



func demo6() {
	var mychan chan int
	mychan = make(chan int, 1)
	mychan <- 1
	for {
		time.Sleep(time.Second)
		number, status := <-mychan
		fmt.Println(number, status)
	}
}
func main() {
	demo6()
}

分析：
    整个程序就只有一个groutine，当channel中没有数据时继续读取，则阻塞，因此DEAD LOCK


func demo3() {
	var myChan chan int
	myChan = make(chan int, 6)
	for i := 0; i < 6; i++ {
		myChan <- i
	}
	fmt.Println("BEGING")
	wg.Add(2)
	go func() {
		temp := 0
		for {
			temp++
			fmt.Println(temp)
			runtime.Gosched()
			time.Sleep(time.Second)
		}
		defer wg.Done()
	}()
    go func() {
		myChan <- 190
		fmt.Println("ADD")
		defer wg.Done()
	}()	
	wg.Wait()
}
分析：
    整个程序两个groutine，第一个groutine在打印了temp的值后，将cpu控制权交给了第二个groutine，第二个groutine向一个已经
    装满数据的channel传入数据，因此第二个groutine阻塞，将cpu控制权交出，但是由于整个程序有两个groutine，因此不会DEAD LOCK
    第一个groutine会继续运行，即使在继续交出CPU控制权，由于第二个groutine一直阻塞，所以即使交出了CPU控制权，第一个groutine
    仍然会继续拿到CPU控制权，继续执行



package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myChan chan int
	myOnce sync.Once
	wg     sync.WaitGroup
)

func read() {
TAG:
	for {
		select {
		case number, status := <-myChan:
			if !status {
				break TAG
			} else {
				fmt.Println(number, status)
				time.Sleep(time.Second)
			}
		default:
			myOnce.Do(func() { close(myChan) })
			break TAG
		}
	}
	defer wg.Done()
}
func demo() {
	myChan = make(chan int, 5)
	for i := 0; i < 5; i++ {
		myChan <- i
	}
	wg.Add(2)
	go read()
	go read()
	wg.Wait()
}
func main() {
	demo()
}
分析：
	select 多路复用的特点，如果在有default的时候，如果所有的case都是阻塞的，则会执行default，因此结合第二种获取channel数据的方式
	可以在接受不到数据发生阻塞时候在default中使用sync.Once关闭channel，然后结束数据获取
	同时在接收数据的case中进行判断channel是否关闭，如果关闭则break。