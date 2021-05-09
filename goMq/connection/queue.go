package connection

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQQueue struct {
	Channel *RabbitMQChannel
	Queue   amqp.Queue
}

// 创建queue
func (c *RabbitMQChannel) CreateQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool,extraInfoMap amqp.Table) *RabbitMQQueue {
	if c == nil {
		return nil
	}
	queue, err := c.Channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, extraInfoMap)
	if err != nil {
		panic(err)
	}
	return &RabbitMQQueue{
		Queue: queue,
	}
}
func (c *RabbitMQChannel) ConnectPassiveQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool) *RabbitMQQueue {
	if c == nil {
		return nil
	}
	queue, err := c.Channel.QueueDeclarePassive(name, durable, autoDelete, exclusive, noWait, nil)
	if err != nil {
		panic(err)
	}
	return &RabbitMQQueue{
		Queue: queue,
	}
}