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

var CLIENT sarama.Client

func initClient() sarama.Client {
	if CLIENT == nil {
		config := sarama.NewConfig()
		config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset
		config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit时间间隔
		CLIENT, err := sarama.NewClient([]string{constant.KAFKA_ADDRESS}, config)
		if err != nil {
			fmt.Printf("Error creating client err: %s\n", err.Error())
			return nil
		} else {
			return CLIENT
		}
	}
	return CLIENT
}

func clientClose(client sarama.Client) {
	if client == nil {
		return
	} else {
		fmt.Println("client close")
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}
}

func getConumserGroupOffsetManager(client sarama.Client, consumerGroupID string) sarama.OffsetManager {
	if client == nil {
		return nil
	} else {
		offsetManager, err := sarama.NewOffsetManagerFromClient(consumerGroupID, client)
		if err != nil {
			return nil
		} else {
			return offsetManager
		}
	}
}
func offsetManagerClose(offsetManager sarama.OffsetManager) {
	if offsetManager == nil {
		return
	} else {
		// 防止自动提交间隔之间的信息被丢掉
		offsetManager.Commit()
		err := offsetManager.Close()
		if err != nil {
			panic(err)
		}
	}
}
func getPartitionManagerOfOffsetManager(offsetManager sarama.OffsetManager, topic string, partition int32) sarama.PartitionOffsetManager {
	if offsetManager == nil {
		return nil
	} else {
		partitionManagerOfOffsetManager, err := offsetManager.ManagePartition(topic, partition)
		if err != nil {
			return nil
		} else {
			return partitionManagerOfOffsetManager
		}
	}
}
func starPartitionManagerOfOffsetManagerErrorChanListen(wg *sync.WaitGroup, partitionManagerOfOffsetManager sarama.PartitionOffsetManager) {
	if partitionManagerOfOffsetManager == nil {
		return
	} else {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range partitionManagerOfOffsetManager.Errors() {
				fmt.Printf("partitionManagerOfOffsetManager  Fail: %+v\n", msg)
			}
		}()
	}
}
func partitionManagerOfOffsetManagerClose(partitionManagerOfOffsetManager sarama.PartitionOffsetManager) {
	if partitionManagerOfOffsetManager == nil {
		return
	} else {
		fmt.Println("partitionManagerOfOffsetManager close")
		partitionManagerOfOffsetManager.AsyncClose()
	}
}
func getConsumerFromClient(client sarama.Client) sarama.Consumer {
	if client == nil {
		return nil
	} else {
		client, err := sarama.NewConsumerFromClient(client)
		if err != nil {
			fmt.Println("create consumer from client error:", err)
			return nil
		} else {
			return client
		}
	}
}

func partitionConsumeWithOffset(ctx context.Context, wg *sync.WaitGroup, partitionConsumer sarama.PartitionConsumer, partitionManagerOfOffsetManager sarama.PartitionOffsetManager) error {
	if partitionConsumer == nil || partitionManagerOfOffsetManager == nil {
		return errors.New("partitionConsumer or partitionManagerOfOffsetManager is nil")
	}
	wg.Add(1)
	defer wg.Done()
	defer partitionManagerOfOffsetManagerClose(partitionManagerOfOffsetManager)
	defer partitionConsumerClose(partitionConsumer)
	partitionConsumeChan := partitionConsumer.Messages()
FORLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			break FORLOOP
		case msg := <-partitionConsumeChan:
			fmt.Printf("Recv Message: [Key:%+v\n Value:%+v,Topic:%s,Partition:%d,Offset:%d]\n", string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
			// 每次消费后都更新一次 offset,这里更新的只是程序内存中的值，需要 commit 之后才能提交到 kafka
			partitionManagerOfOffsetManager.MarkOffset(msg.Offset+1, "has been consumed")
		}
	}
	return nil
}
func ConsumeWithOffsetManager(ctx context.Context, groupID string, topic string) error {
	client := initClient()
	defer clientClose(client)
	// 根据 groupID 来区分不同的 consumer
	//!!!注意: 每次提交的 offset 信息也是和 groupID 关联的
	offsetManager := getConumserGroupOffsetManager(client, groupID)
	defer offsetManagerClose(offsetManager)
	// 每个分区的 offset 也是分别管理的
	partitionManagerOfOffsetManager := getPartitionManagerOfOffsetManager(offsetManager, topic, constant.PARTITION_0)
	consumer := getConsumerFromClient(client)
	defer consumerClose(consumer)
	partitions, err := client.Partitions(topic)
	if err != nil {
		return err
	}
	fmt.Println("the all partitions are: ", partitions)
	partition := partitions[0]
	fmt.Println("use partition is: ", partition)
	allOffset, err := client.GetOffset(topic, partition, -1)
	if err != nil {
		return err
	}
	fmt.Println("all offsets is: ", allOffset)
	nextOffsetToConsume, _ := partitionManagerOfOffsetManager.NextOffset()
	fmt.Println("the nextOffsetToConsume is: ", nextOffsetToConsume)

	var offsetBeginToConsume int64 = 0
	if nextOffsetToConsume == -1 {
		offsetBeginToConsume = 0
	} else {
		offsetBeginToConsume = nextOffsetToConsume
	}
	fmt.Printf("the remain message in partition %d are: %d\n ", partition, allOffset-offsetBeginToConsume)
	partitionConsumer := getPartitionConsumer(consumer, topic, partition, offsetBeginToConsume)
	wg := &sync.WaitGroup{}
	startConsumerErrorChanListen(wg, partitionConsumer)
	starPartitionManagerOfOffsetManagerErrorChanListen(wg, partitionManagerOfOffsetManager)
	partitionConsumeWithOffset(ctx, wg, partitionConsumer, partitionManagerOfOffsetManager)
	wg.Wait()
	return nil
}
