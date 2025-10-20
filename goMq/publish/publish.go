package publish

import (
    "context"

    amqp "github.com/rabbitmq/amqp091-go"

    "mq_demo/connection"
)

// PublishWithContext 在指定的 exchange 和 routing key 上发布一条消息（带上下文）
// mandatory 表示消息必须被路由到队列，immediate 表示消费者必须立即可用，persistent 表示消息持久化
func PublishWithContext(ctx context.Context, channel *connection.RabbitMQChannel, exchange string, queueName string, mandatory bool, immediate bool, persistent bool) *amqp.Publishing {
    if channel == nil {
        panic("无效的RabbitMQChannel")
    }
    msg := "hello,world"
    mode := amqp.Transient
    if persistent {
        mode = amqp.Persistent
    }
    amqpMsg := amqp.Publishing{
        DeliveryMode: mode,
        ContentType:  "text/plain",
        Body:         []byte(msg),
    }
    err := channel.Channel.PublishWithContext(ctx, exchange, queueName, mandatory, immediate, amqpMsg)
    if err != nil {
        panic(err)
    }
    return &amqpMsg
}

// PublishWithDeferredConfirmWithContext 发布消息并返回一个可等待的确认对象（带上下文）
// mandatory 表示消息必须被路由到队列，immediate 表示消费者必须立即可用，persistent 表示消息持久化
func PublishWithDeferredConfirmWithContext(ctx context.Context, channel *connection.RabbitMQChannel, exchange string, routingKey string, mandatory bool, immediate bool, persistent bool) *amqp.DeferredConfirmation {
    if channel == nil || channel.Channel == nil {
        panic("无效的RabbitMQChannel")
    }
    msg := "hello,world"
    mode := amqp.Transient
    if persistent {
        mode = amqp.Persistent
    }
    amqpMsg := amqp.Publishing{
        DeliveryMode: mode,
        ContentType:  "text/plain",
        Body:         []byte(msg),
    }
    deferredConfirmation, err := channel.Channel.PublishWithDeferredConfirmWithContext(ctx, exchange, routingKey, mandatory, immediate, amqpMsg)
    if err != nil {
        panic(err)
    }
    if deferredConfirmation == nil {
        panic("channel的Confirm没有打开")
    }
    return deferredConfirmation
}
