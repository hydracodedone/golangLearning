package consumer

import (
	"context"
	"errors"
	"fmt"
	"go-kafka/constant"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

var CONSUMER sarama.Consumer

func initConsumer() sarama.Consumer {
	if CONSUMER == nil {
		config := sarama.NewConfig()
		config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset
		config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit时间间隔

		CONSUMER, err := sarama.NewConsumer([]string{constant.KAFKA_ADDRESS}, config)
		if err != nil {
			fmt.Printf("Error creating consumer err: %s\n", err.Error())
			return nil
		}
		return CONSUMER
	}
	return CONSUMER
}
func consumerClose(consumer sarama.Consumer) {
	if consumer == nil {
		return
	} else {
		fmt.Println("consumer close")
		err := consumer.Close()
		if err != nil {
			panic(err)
		}
	}
}

func getPartitionConsumer(consumer sarama.Consumer, topic string, partition int32, offset int64) sarama.PartitionConsumer {
	if consumer == nil {
		return nil
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		fmt.Printf("Error creating partition consumer err: %s\n", err.Error())
		return nil
	}
	return partitionConsumer
}
func partitionConsumerClose(partitionConsumer sarama.PartitionConsumer) {
	if partitionConsumer == nil {
		return
	} else {
		fmt.Println("partitionConsumer async close")
		partitionConsumer.AsyncClose()
	}
}
func startConsumerErrorChanListen(wg *sync.WaitGroup, partitionConsumer sarama.PartitionConsumer) {
	if partitionConsumer == nil {
		return
	} else {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range partitionConsumer.Errors() {
				fmt.Printf("Recv Message Fail: %+v\n", msg)
			}
		}()
	}
}
func partitionConsumerConsume(ctx context.Context, wg *sync.WaitGroup, partitionConsumer sarama.PartitionConsumer) {
	wg.Add(1)
	defer wg.Done()
	defer partitionConsumerClose(partitionConsumer)
	consumeChan := partitionConsumer.Messages()
FORLOOP:
	for {
		select {
		case <-ctx.Done():
			break FORLOOP
		case msg := <-consumeChan:
			fmt.Printf("Recv Message: [Key:%+v\n Value:%+v,Topic:%s,Partition:%d,Offset:%d]\n", string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
		}
	}
}

func ConsumeWithAllTopicPartitions(ctx context.Context, topic string) error {
	consumer := initConsumer()
	if consumer == nil {
		err := errors.New("nil consumer")
		return err
	}
	defer consumerClose(consumer)
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	for _, each_partion := range partitions {
		partitionConsumer := getPartitionConsumer(consumer, topic, each_partion, sarama.OffsetOldest)
		if partitionConsumer == nil {
			err := errors.New("nil partition consumer")
			return err
		}
		startConsumerErrorChanListen(wg, partitionConsumer)
		partitionConsumerConsume(ctx, wg, partitionConsumer)
		wg.Wait()
	}
	return nil
}
