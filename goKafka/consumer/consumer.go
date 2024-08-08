package consumer

import (
	"context"
	"fmt"
	"go-kafka/constant"
	"sync"

	"github.com/IBM/sarama"
)

type consumer struct {
	config   *sarama.Config
	consumer sarama.Consumer
}

func GetConsumer() *consumer {
	ap := &consumer{}
	ap.initConsumerConfig()
	ap.initConsumer()
	return ap
}

func (ap *consumer) initConsumerConfig() {
	config := sarama.NewConfig()
	ap.config = config
}

func (ap *consumer) initConsumer() {
	consumer, err := sarama.NewConsumer([]string{constant.KAFKA_ADDRESS}, ap.config)
	if err != nil {
		panic(err)
	} else {
		ap.consumer = consumer
	}
}

func (ap *consumer) closeConsumer() {
	if ap.consumer != nil {
		err := ap.consumer.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("syncConsumer closed")
	}
}

func (ap *consumer) RecvDataWithCancelCtx(ctx context.Context) {
	if ap.consumer == nil {
		panic(constant.ErrKafakaConsusmerNotInitialized)
	} else {
		defer ap.closeConsumer()
		partitionConsumer, err := ap.consumer.ConsumePartition(constant.TOPIC, 0, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		messsageChan := partitionConsumer.Messages()
		go func() {
			<-ctx.Done()
			err := partitionConsumer.Close()
			if err != nil {
				panic(err)
			}
			fmt.Println("partitionConsumer closed")
		}()
		fmt.Println("begin")
		for eachMessage := range messsageChan {
			fmt.Printf("data receiving successfully,message:[%#v],topic:[%s],key:[%v],partition:[%v],offset:[%v]\n", string(eachMessage.Value), eachMessage.Topic, string(eachMessage.Key), 0, eachMessage.Offset)
		}
	}
}

func (ap *consumer) RecvDataWithCancelCtxWithAllPartitions(ctx context.Context) {
	if ap.consumer == nil {
		panic(constant.ErrKafakaConsusmerNotInitialized)
	} else {
		defer ap.closeConsumer()
		partitions, err := ap.consumer.Partitions(constant.TOPIC)
		if err != nil {
			panic(err)
		}
		partitionConsumeWaitGroup := &sync.WaitGroup{}
		for eachPartionNumber := range partitions {
			partitionConsumeWaitGroup.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup, parititionNumber int) {
				defer wg.Done()
				partitionConsumer, err := ap.consumer.ConsumePartition(constant.TOPIC, int32(parititionNumber), sarama.OffsetOldest)
				if err != nil {
					panic(err)
				}
				messsageChan := partitionConsumer.Messages()
				go func() {
					<-ctx.Done()
					err := partitionConsumer.Close()
					if err != nil {
						panic(err)
					}
					fmt.Printf("partitionConsumer closed,partition:[%d]\n", parititionNumber)
				}()
				fmt.Printf("begin to consume topic %s,partition:%d\n", constant.TOPIC, parititionNumber)
				for eachMessage := range messsageChan {
					fmt.Printf("data receiving successfully,message:[%#v],topic:[%s],key:[%v],partition:[%v],offset:[%v]\n", string(eachMessage.Value), eachMessage.Topic, string(eachMessage.Key), eachMessage.Partition, eachMessage.Offset)
					
				}
			}(ctx, partitionConsumeWaitGroup, eachPartionNumber)
		}
		partitionConsumeWaitGroup.Wait()
	}
}
