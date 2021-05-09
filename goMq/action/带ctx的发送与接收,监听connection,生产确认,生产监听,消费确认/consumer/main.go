package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sourcegraph/conc"

	"mq_demo/connection"
	"mq_demo/consume"
)

func main() {
	queueName := "test"
	consumeName := "hydra"
	exclusive := true
	noWait := true
	autoAck := false
	con := connection.GetNewConnection()
	con.ConnectionSetNotifyChannel()
	defer con.CloseConnection()
	channel := con.GetNewChannel()
	defer channel.CloseChannel()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	deliveryChan := consume.GetContextDelivery(ctx, channel, queueName, consumeName, autoAck, exclusive, noWait)
	wg := conc.NewWaitGroup()
	handleFunc := func() {
		for eachDeliveryMsg := range deliveryChan {
			err := channel.Channel.Ack(eachDeliveryMsg.DeliveryTag, false)
			// err := eachDeliveryMsg.Ack(false)
			if err != nil {
				panic("ack fail")
			}
			fmt.Println(string(eachDeliveryMsg.Body))
		}
	}
	connectionNotifyFunc := func() {
	endloopSelct:
		for {
			select {
			case blockNotify := <-con.ConBlockChan:
				fmt.Printf("connection is blocked:[%v]\n", blockNotify)
			case closeNotify := <-con.ConCloseChan:
				fmt.Printf("connection is closed:[%v]\n", closeNotify)
				break endloopSelct
			}
		}
	}
	wg.Go(handleFunc)
	go connectionNotifyFunc()
	wg.Wait()
}
