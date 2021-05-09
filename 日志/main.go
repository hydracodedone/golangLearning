package main

import (
	"fmt"
	"log"
	"os"
)

var myLogger *log.Logger

func init() {
	logFile, err := os.OpenFile("./日志/demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Panicf("Open Or Create Log File Fail:<%s>", err.Error())
	}
	myLogger = log.New(logFile, "[init]", log.Ldate|log.Ltime|log.Llongfile)

}
func logFatalDemo() {
	log.Fatalln("log fatal demo")
}
func logPanicDemo() {
	defer func() {
		errIns := recover()
		if errIns != nil {
			fmt.Printf("The errIns is %s\n", errIns)
		}
	}()
	log.Panic("log panic demo")
}
func logFlagDemo() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("[main.go]")
	log.Println("This is a accurate log ")
}
func logFileDemo() {
	logFile, err := os.OpenFile("./log/demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Panicf("Open Or Create Log File Fail:<%s>", err.Error())
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(logFile)
	log.Println("This is a accurate log write into file")
	defer logFile.Close()
}
func initLoggerDemo() {
	myLogger.Println("This is a accurage log useing default logger")
}
func main() {
	initLoggerDemo()
}
