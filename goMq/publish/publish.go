package publish

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"

	"mq_demo/connection"
)

func PublishWithContext(ctx context.Context, channel *connection.RabbitMQChannel, exchange string, queueName string, mandatory bool, immediat bool, persistent bool) *amqp.Publishing {
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
	err := channel.Channel.PublishWithContext(ctx, exchange, queueName, mandatory, immediat, amqpMsg)
	if err != nil {
		panic(err)
	}
	return &amqpMsg
}

func PublishWithDeferredConfirmWithContext(ctx context.Context, channel *connection.RabbitMQChannel, exchange string, routingKey string, mandatory bool, immediat bool, persistent bool) *amqp.DeferredConfirmation {
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
	deferredConfirmation, err := channel.Channel.PublishWithDeferredConfirmWithContext(ctx, exchange, routingKey, mandatory, immediat, amqpMsg)
	if err != nil {
		panic(err)
	}
	if deferredConfirmation == nil {
		panic("channel的Confirm没有打开")
	}
	return deferredConfirmation
}
