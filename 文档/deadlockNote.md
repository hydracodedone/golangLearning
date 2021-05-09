## channel 数据接收的两种方式

info <-channel 

	如果channel没有关闭且channel中没有数据，则会阻塞该语句所在的channel
	如果channel关闭，如果有数据，则返回数据，没有数据，则返回0

info,status <-channel

	如果channel关闭后，如果channel中的数据被读取完毕后，会返回 0，false
	关闭channel后，如果channel中的数据没有被读取完，会返回数据，ture
	没有关闭，有数据，会返回数据，true
	没有关闭，没有数据，会阻塞

## deadlock

所有的groutine都等待(阻塞)则发生死锁

### 阻塞

装满以后如果继续输入则会阻塞该输入所在的groutine，并将cpu使用权交出，sleep
为空如果继续输出则会发生阻塞该输出所在的groutine，并将cpu使用权交出，sleep


### 常见死锁分析
```Go
package main

import "fmt"

func main() {
	var myChan chan int
	myChan = make(chan int, 6)
	for i := 0; i < 6; i++ {
		myChan <- i
	}
	myChan <- 90
	fmt.Println("SS")
}
```

分析：
在main groutine 中，由于myChan已经装满，因此，myChan <- 90 会阻塞当前 的main对应的groutine,所有的groutine都sleep，因此死锁

```Go
package main

import (
	"fmt"
	"time"
)

func main() {
	var mychan chan int
	mychan = make(chan int, 1)
	mychan <- 1
	for {
		time.Sleep(time.Second)
		number, status := <-mychan
		fmt.Println(number, status)
	}
}
```
分析：
整个程序就只有一个groutine，当channel中没有数据时继续读取，则阻塞，因此DEAD LOCK

```Go
package main

import "fmt"

func main() {
	var myChan chan int
	myChan = make(chan int, 6)
	for i := 0; i < 60; i++ {
		myChan <- i
	}
	go func() {
		for recv := range myChan {
			fmt.Println(recv)
		}
	}()
}
```
分析：
即使有尝试开启从chan中取的goroutine,但是阻塞发生时候并没有执行到开启goroutine的代码段,因此依然会发生死锁
