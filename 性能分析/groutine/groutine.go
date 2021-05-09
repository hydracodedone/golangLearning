package groutine

import "time"

func g_cal() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second * 2)
	}
}
func GroutineDemo() {
	for i := 0; i < 1000; i++ {
		go g_cal()
	}
	<-time.After(time.Minute * 10)
}
