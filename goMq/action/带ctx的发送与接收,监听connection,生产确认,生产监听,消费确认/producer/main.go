package main

import (
	"context"
	"fmt"
	"time"

	"mq_demo/connection"
	"mq_demo/publish"
)

func main() {
	queueName := "test"
	durable := true     //持久化
	autoDelete := false //不自动删除
	exclusive := false
	noWait := true //不需要等待

	con := connection.GetNewConnection()
	defer con.CloseConnection()
	channel := con.GetNewConfirmChannel()
	defer channel.CloseChannel()
	channel.CreateQueue(queueName, durable, autoDelete, exclusive, noWait, nil)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	exchangeName := "" //默认交换机
	mandatory := true  //要求能够被路由到某个queue
	immedate := false  //要求由消费者可以不在线
	persistent := true //消息要求持久化
	//发布消息确认

	//发布
	deferredConfirmedWithContext := publish.PublishWithDeferredConfirmWithContext(ctx, channel, exchangeName, queueName, mandatory, immedate, persistent)
	acked, err := deferredConfirmedWithContext.WaitContext(ctx)
	if err != nil {
		fmt.Println("deferredConfirmedWithContext 监听超时")
		panic(err)
	}
	if acked {
		fmt.Printf("deferredConfirmedWithContext 消息已经发布成功,delivery Tag:[%v]\n", deferredConfirmedWithContext.DeliveryTag)
	} else {
		fmt.Printf("deferredConfirmedWithContext 消息已经发布失败,delivery Tag:[%v]\n", deferredConfirmedWithContext.DeliveryTag)
	}
	sig := make(chan struct{}, 1)
	go func(c chan struct{}) {
	endloopSelct:
		for {
			select {
			case ackReturn, ok := <-channel.NotifyReturtn:
				if !ok {
					fmt.Println("return 监听通道已经关闭")
					break endloopSelct
				}
				fmt.Printf("监听到return消息:Return Message[%v]\n", ackReturn)
			case ackNotify, ok := <-channel.NotifyAckChan:
				if !ok {
					fmt.Println("ack 监听通道已经关闭")
					break endloopSelct
				}
				fmt.Printf("监听到ack消息:ACK[%d]\n", ackNotify)
			case nackNotify, ok := <-channel.NotifyNackChan:
				if !ok {
					fmt.Println("nack 监听通道已经关闭")
					break endloopSelct
				}
				fmt.Printf("监听到nack消息:NACK[%d]\n", nackNotify)
			//确认收到消息循环队列
			case publishNotify, ok := <-channel.NotifyPublishChan:
				if !ok {
					fmt.Println("publish 监听通道已经关闭")
					break endloopSelct
				}
				fmt.Printf("监听到publish消息DeliverTag:[%x]已经发送,ACK[%t]\n", publishNotify.DeliveryTag, publishNotify.Ack)
			case <-sig:
				fmt.Println("publish 监听进程收到退出信号sig")
				break endloopSelct
			}
		}
		fmt.Println("publish 监听进程退出")
	}(sig)
	time.Sleep(time.Second)
	sig <- struct{}{}
	fmt.Println("main 函数退出")
}
