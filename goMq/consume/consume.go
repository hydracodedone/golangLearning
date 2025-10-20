package consume

import (
    "context"
    "mq_demo/connection"

    amqp "github.com/rabbitmq/amqp091-go"
)

// GetDelivery 从指定队列获取消费通道（不带上下文）
// autoAck 表示是否自动确认，exclusive 表示是否排他，noWait 表示是否不等待服务器响应
func GetDelivery(channel *connection.RabbitMQChannel, queueName string, consumerName string, autoAck bool, exclusive bool, noWait bool) <-chan amqp.Delivery {
    if channel == nil {
        panic("无效的RabbitMQChannel")
    }
    deliveryChan, err := channel.Channel.Consume(queueName, consumerName, autoAck, exclusive, true, noWait, nil)
    if err != nil {
        panic("获取消费通道失败")
    }
    return deliveryChan
}

// GetContextDelivery 从指定队列获取带上下文的消费通道
// autoAck 表示是否自动确认，exclusive 表示是否排他，noWait 表示是否不等待服务器响应
func GetContextDelivery(ctx context.Context, channel *connection.RabbitMQChannel, queueName string, consumerName string, autoAck bool, exclusive bool, noWait bool) <-chan amqp.Delivery {
    if channel == nil {
        panic("无效的RabbitMQChannel")
    }
    deliveryChan, err := channel.Channel.ConsumeWithContext(ctx, queueName, consumerName, autoAck, exclusive, true, noWait, nil)
    if err != nil {
        panic("获取消费通道失败")
    }
    return deliveryChan
}
