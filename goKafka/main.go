package main

import (
	"context"
	"fmt"
	"go-kafka/asyncProducer"
	"go-kafka/consumer"
	"go-kafka/syncProducer"

	"time"
)

func AsyncProduce() {
	ap := asyncProducer.GetAsyncProducer()
	defer ap.CloseAsyncProducer()
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello World %d", i)
		ap.SendData(msg)
	}
}
func SyncProduce() {
	sp := syncProducer.GetSyncProducer()
	defer sp.ClosesyncProducer()
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello World %d", i)
		err := sp.SendData(msg)
		if err != nil {
			panic(err)
		}
	}
}

func StandaloneConsume() {
	sc := consumer.GetConsumer()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		timer := time.NewTimer(10 * time.Second)
		<-timer.C
		cancelFunc()
	}()
	sc.RecvDataWithCancelCtx(ctx)
}
func ConsumeAllPartitions() {
	sc := consumer.GetConsumer()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		timer := time.NewTimer(10 * time.Second)
		<-timer.C
		cancelFunc()
	}()
	sc.RecvDataWithCancelCtxWithAllPartitions(ctx)
}
func ConsumeAllPartitionsWithOffset() {
	sc := consumer.GetOffsetConsumer()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		timer := time.NewTimer(4 * time.Second)
		<-timer.C
		cancelFunc()
	}()
	sc.RecvDataWithCancelCtxWithAllPartitions(ctx)
}
func ConsumeWithConsumeGroup(name string) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		timer := time.NewTimer(20 * time.Second)
		<-timer.C
		cancelFunc()
	}()
	kc := consumer.GetkafkaConsumeInstance()
	kc.ConsumeByConsumerGroup(ctx, name)
}

func Rebalance() {
	go ConsumeWithConsumeGroup("consumer1")
	time.Sleep(6 * time.Second)
	ConsumeWithConsumeGroup("consumer2")
	time.Sleep(15 * time.Second)
}
func main() {
	AsyncProduce()
	// ConsumeWithConsumeGroup("consumer1")
	// ConsumeAllPartitionsWithOffset()
	// Rebalance()
	// ConsumeAllPartitions()
}
