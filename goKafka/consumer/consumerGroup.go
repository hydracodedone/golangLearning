package consumer

import (
	"context"
	"fmt"
	"go-kafka/constant"
	"sync"

	"github.com/IBM/sarama"
)

type kafkaConsumeInstance struct {
	config          *sarama.Config
	client          sarama.Client
	consumerGroup   sarama.ConsumerGroup
	consumerGroupID string
	wg              *sync.WaitGroup
}

type consumerGroupHandler struct {
	consumerGroupID string
	consumerName    string
}

func (cgh *consumerGroupHandler) initConsumerGroupHandler(name string) {
	cgh.consumerGroupID = constant.CONSUMER_GROUP
	cgh.consumerName = name

}
func (cgh *consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("consumerGroupHandler %s setup\n", cgh.consumerName)
	fmt.Println(session.Claims())
	return nil
}
func (cgh *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	fmt.Printf("consumerGroupHandler %s cleanup\n", cgh.consumerName)
	return nil
}
func (cgh *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer sess.Commit()
	fmt.Printf("consumerGroupHandler %s begin consume partition %d\n", cgh.consumerName, claim.Partition())
	autoCommitInterval := 5
	initialCommit := 0
	for msg := range claim.Messages() {
		fmt.Printf("consumerGroupHandler %s consume topic:%q partition:%d offset:%d\n", cgh.consumerName, msg.Topic, msg.Partition, msg.Offset)
		// 标记消息已被消费 内部会更新 consumer offset
		sess.MarkMessage(msg, fmt.Sprintf("consume by %s", cgh.consumerName))
		initialCommit += 1
		if initialCommit%autoCommitInterval == 0 {
			sess.Commit()
			initialCommit = 0
		}
		sess.Commit()
	}
	return nil
}

func GetkafkaConsumeInstance() *kafkaConsumeInstance {
	kc := &kafkaConsumeInstance{}
	kc.initConfig()
	client, err := sarama.NewClient([]string{constant.KAFKA_ADDRESS}, kc.config)
	if err != nil {
		panic(err)
	}
	kc.client = client
	kc.consumerGroupID = constant.CONSUMER_GROUP
	consumerGroup, err := sarama.NewConsumerGroupFromClient(kc.consumerGroupID, client)
	if err != nil {
		panic(err)
	}
	kc.consumerGroup = consumerGroup
	kc.wg = &sync.WaitGroup{}
	return kc
}

func (kc *kafkaConsumeInstance) initConfig() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.Retry.Max = 3
	kc.config = config
}

func (kc *kafkaConsumeInstance) Close() {
	if kc.consumerGroup != nil {
		err := kc.consumerGroup.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("consumerGroup closed")
	}
	if kc.client != nil {
		err := kc.client.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("client closed")

	}
}
func (kc *kafkaConsumeInstance) ConsumeByConsumerGroup(ctx context.Context, name string) {
	defer kc.Close()
	kc.wg.Add(1)
	go func() {
		defer kc.wg.Done()
		handler := consumerGroupHandler{}
		handler.initConsumerGroupHandler(name)
		for {
			err := kc.consumerGroup.Consume(ctx, []string{constant.TOPIC}, &handler)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	kc.wg.Wait()
}
