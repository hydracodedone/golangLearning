package asyncProducer

import (
	"context"
	"fmt"
	"go-kafka/constant"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

type AsyncProducer struct {
	config                  *sarama.Config
	ctx                     context.Context
	signal                  chan struct{}
	signalForSuccessChannel chan struct{}
	signalForFailChannel    chan struct{}
	asyncProducer           sarama.AsyncProducer
	cancelFunc              context.CancelFunc
	Error                   error
}

func GetAsyncProducer() *AsyncProducer {
	ap := &AsyncProducer{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	ap.ctx = ctx
	ap.cancelFunc = cancelFunc
	ap.initAsyncProducerConfig()
	ap.initAsyncProducer()
	ap.signal = make(chan struct{})
	ap.signalForSuccessChannel = make(chan struct{})
	ap.signalForFailChannel = make(chan struct{})
	return ap
}

func (ap *AsyncProducer) initAsyncProducerConfig() {
	config := sarama.NewConfig()
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewManualPartitioner // 新选出一个partition,对于单节点的kafka貌似不行,即使手动指定了不同分区,同一个topic也会将手动指定的不同分区的放在一个partition
	config.Producer.Retry.Max = 3
	// config.Producer.Compression = sarama.CompressionSnappy //开启后 kafka-map无法正常显示
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	ap.config = config
}
func (ap *AsyncProducer) initAsyncProducer() {
	asyncProducer, err := sarama.NewAsyncProducer([]string{constant.KAFKA_ADDRESS}, ap.config)
	if err != nil {
		ap.Error = err
	} else {
		ap.asyncProducer = asyncProducer
		ap.listenSuccessAndFailChannel()
	}
}
func (ap *AsyncProducer) GetError() error {
	return ap.Error
}
func (ap *AsyncProducer) CloseAsyncProducer() {
	if ap.asyncProducer != nil {
		func() {
			ap.asyncProducer.AsyncClose()
			ap.cancelFunc()
			close(ap.signal)
		}()
		<-ap.signalForSuccessChannel
		<-ap.signalForFailChannel
	}
}

func (ap *AsyncProducer) listenSuccessAndFailChannel() {
	if ap.asyncProducer == nil {
		return
	} else {
		go func() {
			for msg := range ap.asyncProducer.Successes() {
				fmt.Printf("Send Message Success: %+v\n", msg)
			}
			fmt.Println("async producer success channel closed")
			<-ap.signal
			close(ap.signalForSuccessChannel)
		}()
		go func() {
			for msg := range ap.asyncProducer.Errors() {
				fmt.Printf("Send Message Error: %+v\n", msg)
			}
			fmt.Println("async producer error channel closed")
			<-ap.signal
			close(ap.signalForFailChannel)
		}()
	}
}
func (ap *AsyncProducer) generateMessage(topic string, key string, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.StringEncoder(message)
	return msg
}

func (ap *AsyncProducer) SendData(message string) error {
	if ap.asyncProducer == nil {
		return constant.ErrKafakaProducerNotInitialized
	} else {
		timeOutCtx, cancelFunc := context.WithTimeout(ap.ctx, time.Second)
		defer cancelFunc()
		msg := ap.generateMessage(constant.TOPIC, constant.KEY, message)
		select {
		case <-timeOutCtx.Done():
			return constant.ErrKafakaProducerSendMessageTimeout
		case <-ap.ctx.Done():
			return constant.ErrKafakaProducerHasClosed
		case ap.asyncProducer.Input() <- msg:
			fmt.Printf("async producer send message to channel:%v\n", msg)
			return nil
		}
	}
}
