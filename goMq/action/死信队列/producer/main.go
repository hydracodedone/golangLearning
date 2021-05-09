package main

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"mq_demo/connection"
	"mq_demo/publish"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	//交换机参数说明
	deadExchageName := "dead_exchange"
	normalExchangeName := "normal_exchange"
	deadRoutingKey := "dead_key"
	normalMessageMaxLength := 100
	normalRoutingKey := "dead_key"
	exchangeType := "direct"
	durable := true
	autoDelete := false
	noWait := true
	deadExchangeInternal := true
	normalExchangeInternal := false
	//队列参数说明
	deadQueueName := "dead_queue"
	normalQueueName := "normal_queue"
	queueDurable := true
	queueAutoDelete := false
	queueExclusive := false
	queueNoWait := true
	ttl := 30000 //3s

	bindNoWait := false
	//publish参数说明
	mandatory := true  //要求能够被路由到某个queue
	immedate := false  //要求由消费者可以不在线
	persistent := true //消息要求持久化
	con := connection.GetNewConnection()
	defer con.CloseConnection()
	channel := con.GetNewConfirmChannel()
	defer channel.CloseChannel()
	//创建发布交换机
	channel.ExchageDeclare(normalExchangeName, exchangeType, durable, autoDelete, normalExchangeInternal, noWait)
	//创建死信交换机
	channel.ExchageDeclare(deadExchageName, exchangeType, durable, autoDelete, deadExchangeInternal, noWait)
	//创建队列
	//死信配置
	extraInfoMap := amqp.Table{
		"x-message-ttl":             ttl,
		"x-max-length":              normalMessageMaxLength,
		"x-dead-letter-exchange":    deadExchageName,
		"x-dead-letter-routing-key": deadRoutingKey,
	}
	//配置死信信息
	channel.CreateQueue(normalQueueName, queueDurable, queueAutoDelete, queueExclusive, queueNoWait, extraInfoMap)
	//创建死信队列
	channel.CreateQueue(deadQueueName, queueDurable, queueAutoDelete, queueExclusive, queueNoWait, nil)
	//绑定
	channel.ExchangeBindQueue(normalQueueName, normalRoutingKey, normalExchangeName, bindNoWait, nil)
	channel.ExchangeBindQueue(deadQueueName, deadRoutingKey, deadExchageName, bindNoWait, nil)

	//发布消息
	deferredConfirmedWithContext := publish.PublishWithDeferredConfirmWithContext(ctx, channel, normalExchangeName, normalRoutingKey, mandatory, immedate, persistent)
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
