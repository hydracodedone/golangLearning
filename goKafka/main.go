package main

import (
	"context"
	"fmt"
	"go-kafka/asyncProducer"
	"go-kafka/constant"
	"go-kafka/consumer"

	"time"
)

func AsyncProduce() {
	ctx, cacelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cacelFunc()
	topic := constant.TOPIC
	key := constant.KEY
	values := make([]string, 0, 50)
	for i := 0; i < 10; i++ {
		values = append(values, fmt.Sprintf("message-%d", i))
	}
	err := asyncProducer.AsyncProduce(ctx, values, topic, key)
	if err != nil {
		panic(err)
	}
}
func Consume() {
	ctx, cacelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cacelFunc()
	topic := constant.TOPIC
	err := consumer.ConsumeWithAllTopicPartitions(ctx, topic)
	if err != nil {
		panic(err)
	}
}
func ConsumeWithOffSet() {
	ctx, cacelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cacelFunc()
	topic := constant.TOPIC
	group_id := constant.CONSUMER_GROUP
	err := consumer.ConsumeWithOffsetManager(ctx, group_id, topic)
	if err != nil {
		panic(err)
	}
}
func main() {
	// AsyncProduce()
	ConsumeWithOffSet()
}
