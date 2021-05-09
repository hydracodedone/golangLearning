package consume

import (
	"context"
	"mq_demo/connection"

	amqp "github.com/rabbitmq/amqp091-go"
)

func GetDelivery(channel *connection.RabbitMQChannel, queueName string, consumerName string, auotAck bool, exclusive bool, noWait bool) <-chan amqp.Delivery {
	if channel == nil {
		panic("无效的RabbitMQChannel")
	}
	diliveryChan, err := channel.Channel.Consume(queueName, consumerName, auotAck, exclusive, true, noWait, nil)
	if err != nil {
		panic("获取消费通道失败")
	}
	return diliveryChan
}

func GetContextDelivery(ctx context.Context, channel *connection.RabbitMQChannel, queueName string, consumerName string, auotAck bool, exclusive bool, noWait bool) <-chan amqp.Delivery {
	if channel == nil {
		panic("无效的RabbitMQChannel")
	}
	diliveryChan, err := channel.Channel.ConsumeWithContext(ctx, queueName, consumerName, auotAck, exclusive, true, noWait, nil)
	if err != nil {
		panic("获取消费通道失败")
	}
	return diliveryChan
}
