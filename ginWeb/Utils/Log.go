package Utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerTest() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	writerFile, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	writerStdOut := os.Stdout
	logrus.SetOutput(io.MultiWriter(writerFile, writerStdOut))
	logrus.Warn("warn message")
	loggerEntry := logrus.WithFields(logrus.Fields{"currentTime": time.Now().UTC().Format(time.RFC1123)})
	loggerEntry.Warn("war message with extra info")
}

/*
type Hook interface {
Levels() []Level
Fire(*Entry) error
}
*/
type MyLogHook struct{}

func (m MyLogHook) Levels() []logrus.Level {
	focouseLevel := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
	return focouseLevel
}
func (m MyLogHook) Fire(entry *logrus.Entry) error {
	writerFile, err := os.OpenFile("./log.log.wf", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	entry.Logger.Out = writerFile
	return nil
}

type MyLogHook2 struct{}

func (m MyLogHook2) Levels() []logrus.Level {
	focouseLevel := []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
	}
	return focouseLevel
}
func (m MyLogHook2) Fire(entry *logrus.Entry) error {
	writerFile, err := os.OpenFile("./log.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	entry.Logger.Out = writerFile
	return nil
}

type MycustomFormatter struct {
}

func (m MycustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	logTime := entry.Time.Format(time.RFC3339)
	logLevel := entry.Level
	logMessage := entry.Message
	var message string
	if entry.HasCaller() {
		formatString := "logFile:[%s] logFunc:[%s] logLine:[%d] logTime:[%s] logLevel:[%s] logMsg: %s\n"
		callerFileName := entry.Caller.File
		callerLine := entry.Caller.Line
		callerFunction := entry.Caller.Function
		message = fmt.Sprintf(formatString, callerFileName, callerFunction, callerLine, logTime, logLevel, logMessage)
	} else {
		formatString := "logTime:[%s] logLevel:[%s] logMsg: %s\n"
		message = fmt.Sprintf(formatString, logTime, logLevel, logMessage)
	}
	b.WriteString(message)
	return b.Bytes(), nil
}
func LoggerHookTest() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&MycustomFormatter{})
	logrus.SetReportCaller(true)
	logrus.AddHook(&MyLogHook{})
	logrus.AddHook(&MyLogHook2{})
}
