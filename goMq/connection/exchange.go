package connection

import amqp "github.com/rabbitmq/amqp091-go"

// ExchangeDeclare 声明交换机
func (c *RabbitMQChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool) {
    err := c.Channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, nil)
    if err != nil {
        panic(err)
    }
}

// ConnectPassiveExchange 被动连接到已存在的交换机
func (c *RabbitMQChannel) ConnectPassiveExchange(name, kind string, durable, autoDelete, internal, noWait bool) {
    err := c.Channel.ExchangeDeclarePassive(name, kind, durable, autoDelete, internal, noWait, nil)
    if err != nil {
        panic(err)
    }
}

// ExchangeBindQueue 将队列绑定到指定交换机
func (c *RabbitMQChannel) ExchangeBindQueue(queueName string, routingKey string, exchangeName string, noWait bool, extraInfoMap amqp.Table) {
    err := c.Channel.QueueBind(queueName, routingKey, exchangeName, noWait, extraInfoMap)
    if err != nil {
        panic(err)
    }
}
