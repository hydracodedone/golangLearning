package crontab

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/toolbox"
)

func InitTask() {
	task := toolbox.NewTask("firstTask", "0/30 * * * * *", myTask)
	toolbox.AddTask("firstTask", task)
}
func myTask() error {
	fmt.Println("ticker is running")
	time.Sleep(5 * time.Second)
	fmt.Println("ticker is done")
	return nil
}
