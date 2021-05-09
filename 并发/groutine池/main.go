package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TaskFunc func()
type Task struct {
	TaskFunction TaskFunc
}
type Pool struct {
	TaskChan     chan *Task
	WorkerNumber int
	TaskCapacity int
}

func NewTask(functionInstance TaskFunc) *Task {
	return &Task{
		TaskFunction: functionInstance,
	}
}
func (taskInstance *Task) ExecuteTask() {
	taskInstance.TaskFunction()
}

func NewPool(workerNum int, TaskCapacity int) *Pool {
	pool := Pool{
		TaskChan:     make(chan *Task, TaskCapacity),
		WorkerNumber: workerNum,
		TaskCapacity: TaskCapacity,
	}
	return &pool
}
func (pool *Pool) Work(number int) {
	timerTemp := time.NewTimer(time.Second * 4)
tag:
	for {
		select {
		case task := <-pool.TaskChan:
			task.ExecuteTask()
			timerTemp.Reset(time.Second * 4)
		case <-timerTemp.C:
			timerTemp.Stop()
			fmt.Printf("Worker %d work Close\n", number)
			break tag
		}
	}
	defer wg.Done()
}
func (pool *Pool) Run() {
	for i := 0; i < pool.WorkerNumber; i++ {
		go pool.Work(i)
	}
}

func MyTask() {
	fmt.Printf("Hello,World,Time Is <%v>\n", time.Now())
}
func main() {
	poolInstance := NewPool(2, 10)
	wg.Add(poolInstance.WorkerNumber + 1)
	poolInstance.Run()
	go func() {
		for i := 0; i < 20; i++ {
			taskInstance := NewTask(MyTask)
			poolInstance.TaskChan <- taskInstance
		}
		defer wg.Done()
	}()
	wg.Wait()
}
