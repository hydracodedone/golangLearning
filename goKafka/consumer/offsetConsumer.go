package consumer

import (
	"context"
	"fmt"
	"go-kafka/constant"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

type offSetConsumer struct {
	config               *sarama.Config
	consumer             sarama.Consumer
	offsetManager        sarama.OffsetManager
	client               sarama.Client
	partitionManagerMap  map[int32]sarama.PartitionOffsetManager
	partitionConsumerMap map[int32]sarama.PartitionConsumer
}

func GetOffsetConsumer() *offSetConsumer {
	ap := &offSetConsumer{}
	ap.partitionManagerMap = make(map[int32]sarama.PartitionOffsetManager)
	ap.partitionConsumerMap = make(map[int32]sarama.PartitionConsumer)
	ap.initSyncOffsetConsumerConfig()
	ap.initSyncOffsetConsumer()
	return ap
}

func (ap *offSetConsumer) initSyncOffsetConsumerConfig() {
	config := sarama.NewConfig()
	ap.config = config
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	ap.config.Consumer.Offsets.AutoCommit.Enable = true //自动提交
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
}
func (ap *offSetConsumer) initSyncOffsetConsumer() {
	client, err := sarama.NewClient([]string{constant.KAFKA_ADDRESS}, ap.config)
	if err != nil {
		panic(err)
	}
	ap.client = client
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	ap.consumer = consumer
	offsetManager, err := sarama.NewOffsetManagerFromClient(constant.CONSUMER_GROUP, client)
	if err != nil {
		panic(err)
	}
	ap.offsetManager = offsetManager
	parititionNumber, err := ap.client.Partitions(constant.TOPIC)
	if err != nil {
		panic(err)
	}
	for eachPartionNumber := range parititionNumber {
		parititionOffsetManager, err := ap.offsetManager.ManagePartition(constant.TOPIC, int32(eachPartionNumber))
		if err != nil {
			panic(err)
		}
		ap.partitionManagerMap[int32(eachPartionNumber)] = parititionOffsetManager
		nextOffset, _ := parititionOffsetManager.NextOffset()
		var offset int64 = 0
		if nextOffset == -1 {
			offset = 0
		} else {
			offset = nextOffset
		}
		partitionConsumer, err := ap.consumer.ConsumePartition(constant.TOPIC, int32(eachPartionNumber), offset)
		if err != nil {
			panic(err)
		}
		ap.partitionConsumerMap[int32(eachPartionNumber)] = partitionConsumer
	}
}

func (ap *offSetConsumer) closeAll() {
	fmt.Println("begin to close")
	for parititionNumber, eachPartitionConsumer := range ap.partitionConsumerMap {
		err := eachPartitionConsumer.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("eachPartitionConsumer:[topic:%s partitionNumber:%d] closed\n", constant.TOPIC, parititionNumber)
	}
	fmt.Println(ap.partitionManagerMap)
	for parititionNumber, eachPartionOffsetManager := range ap.partitionManagerMap {
		err := eachPartionOffsetManager.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("partitionOffsetManager:[topic:%s partitionNumber:%d] closed\n", constant.TOPIC, parititionNumber)
	}

	if ap.offsetManager != nil {
		err := ap.offsetManager.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("offsetManager closed")
	}
	if ap.consumer != nil {
		err := ap.consumer.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("consumer closed")
	}
	if ap.client != nil {
		err := ap.client.Close()
		if err != nil {
			fmt.Println("client closed")
		}
	}
	fmt.Println("end close")
}

func (ap *offSetConsumer) RecvDataWithCancelCtxWithAllPartitions(ctx context.Context) {
	ap.offsetManager.Commit()
	if ap.consumer == nil {
		panic(constant.ErrKafakaConsusmerNotInitialized)
	} else {
		go func() {
			<-ctx.Done()
			ap.closeAll()
		}()
		partitionConsumeWaitGroup := &sync.WaitGroup{}
		for eachPartionNumber, parititionOffsetManager := range ap.partitionManagerMap {
			partitionConsumeWaitGroup.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup, parititionNumber int32, parititionOffsetManager sarama.PartitionOffsetManager) {
				defer wg.Done()
				partitionConsumer := ap.partitionConsumerMap[parititionNumber]
				messsageChan := partitionConsumer.Messages()
				fmt.Printf("begin to consume topic %s,partition:%d\n", constant.TOPIC, parititionNumber)
				for eachMessage := range messsageChan {
					parititionOffsetManager.MarkOffset(eachMessage.Offset+1, "modified metadata")
					// ap.offsetManager.Commit()
					fmt.Printf("data receiving successfully,message:[%#v],topic:[%s],key:[%v],partition:[%v],offset:[%v]\n", string(eachMessage.Value), eachMessage.Topic, string(eachMessage.Key), eachMessage.Partition, eachMessage.Offset)
				}
			}(ctx, partitionConsumeWaitGroup, eachPartionNumber, parititionOffsetManager)
		}
		partitionConsumeWaitGroup.Wait()
	}
}