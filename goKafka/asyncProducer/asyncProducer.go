package asyncProducer

import (
	"context"
	"errors"
	"fmt"
	"go-kafka/constant"
	"sync"

	"github.com/IBM/sarama"
)

var ASYNC_PRODUCER sarama.AsyncProducer

func getAsyncProducer() sarama.AsyncProducer {
	if ASYNC_PRODUCER == nil {
		config := sarama.NewConfig()
		config.Producer.Idempotent = true
		config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
		config.Producer.Partitioner = sarama.NewManualPartitioner // 新选出一个partition,对于单节点的kafka貌似不行,即使手动指定了不同分区,同一个topic也会将手动指定的不同分区的放在一个partition
		config.Producer.Retry.Max = 3
		// config.Producer.Compression = sarama.CompressionSnappy //开启后 kafka-map无法正常显示

		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true

		ASYNC_PRODUCER, err := sarama.NewAsyncProducer([]string{constant.KAFKA_ADDRESS}, config)
		if err != nil {
			fmt.Printf("Error creating async producer err: %s\n", err.Error())
			return nil
		}
		return ASYNC_PRODUCER
	}
	return ASYNC_PRODUCER
}
func producerAsyncClose(asyncProducer sarama.AsyncProducer) {
	if asyncProducer != nil {
		fmt.Println("Aysnc Input Close")
		asyncProducer.AsyncClose()
	}
}
func startSuccessAndFailChanListen(wg *sync.WaitGroup, asyncProducer sarama.AsyncProducer) {
	if asyncProducer == nil {
		return
	} else {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range asyncProducer.Successes() {
				fmt.Printf("Send Message Success: %+v\n", msg)
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range asyncProducer.Errors() {
				fmt.Println("Send Message Error:", msg.Err)
				fmt.Println("Send Message Error Message:", msg.Msg)
			}
		}()
	}
}
func asyncSendData(ctx context.Context, wg *sync.WaitGroup, asyncProducer sarama.AsyncProducer, values []string, topic string, key string) (err error) {
	wg.Add(1)
	defer wg.Done()
	defer producerAsyncClose(asyncProducer)
	index := 0
	sendChan := asyncProducer.Input()
	msgLen := len(values)
FORLOOP:
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Key = sarama.StringEncoder(key)
		msg.Value = sarama.StringEncoder(values[index])
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			break FORLOOP
		case sendChan <- msg:
			fmt.Printf("Send Message: %+v\n", msg)
			index += 1
			if index >= msgLen {
				break FORLOOP
			}
		}
	}
	return nil
}
func AsyncProduce(ctx context.Context, values []string, topic string, key string) error {
	asyncProducer := getAsyncProducer()
	if asyncProducer == nil {
		err := errors.New("nil async producer")
		return err
	} else {
		wg := &sync.WaitGroup{}
		startSuccessAndFailChanListen(wg, asyncProducer)
		asyncSendData(ctx, wg, asyncProducer, values, topic, key)
		wg.Wait()
	}
	return nil
}
