package syncProducer

import (
	"fmt"
	"go-kafka/constant"

	"github.com/IBM/sarama"
)

type syncProducer struct {
	config       *sarama.Config
	syncProducer sarama.SyncProducer
}

func GetSyncProducer() *syncProducer {
	ap := &syncProducer{}
	ap.initsyncProducerConfig()
	ap.initsyncProducer()
	return ap
}

func (ap *syncProducer) initsyncProducerConfig() {
	config := sarama.NewConfig()
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition,对于单节点的kafka貌似不行,即使手动指定了不同分区,同一个topic也会将手动指定的不同分区的放在一个partition
	config.Producer.Retry.Max = 3
	config.Producer.Timeout = 1
	ap.config = config
}
func (ap *syncProducer) initsyncProducer() {
	syncProducer, err := sarama.NewSyncProducer([]string{constant.KAFKA_ADDRESS}, ap.config)
	if err != nil {
		panic(err)
	} else {
		ap.syncProducer = syncProducer
	}
}

func (ap *syncProducer) ClosesyncProducer() {
	if ap.syncProducer != nil {
		err := ap.syncProducer.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("syncProducer Closed")
	}
}

func (ap *syncProducer) generateMessage(topic string, key string, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.StringEncoder(message)
	return msg
}

func (ap *syncProducer) SendData(message string) error {
	if ap.syncProducer == nil {
		return constant.ErrKafakaProducerNotInitialized
	} else {
		msg := ap.generateMessage(constant.TOPIC, constant.KEY, message)
		partition, offset, err := ap.syncProducer.SendMessage(msg)
		if err != nil {
			return err
		} else {
			fmt.Printf("data sending successfully,message:[%s],topic:[%s],key:[%v],partition:[%v],offset:[%v]\n", message, constant.TOPIC, constant.KEY, partition, offset)
			return nil
		}
	}
}
