package connection

import amqp "github.com/rabbitmq/amqp091-go"

func (c *RabbitMQChannel) ExchageDeclare(name, kind string, durable, autoDelete, internal, noWait bool) {
	err := c.Channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, nil)
	if err != nil {
		panic(err)
	}
}

func (c *RabbitMQChannel) ConnectPassiveExchange(name, kind string, durable, autoDelete, internal, noWait bool) {
	err := c.Channel.ExchangeDeclarePassive(name, kind, durable, autoDelete, internal, noWait, nil)
	if err != nil {
		panic(err)
	}
}
func (c *RabbitMQChannel) ExchangeBindQueue(queueName string, routingKey string, exchangeName string, noWait bool, extraInfoMap amqp.Table) {
	err := c.Channel.QueueBind(queueName, routingKey, exchangeName, noWait, extraInfoMap)
	if err != nil {
		panic(err)
	}
}
