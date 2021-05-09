package main

import "log"
import "os"
import "fmt"

func init() {
	logFile, err := os.OpenFile("./logFile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	log.SetPrefix("[loggerDemo]")
	log.SetOutput(logFile)

}
func main() {
	log.Printf("this is a test")
}
