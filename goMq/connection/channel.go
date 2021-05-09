package connection

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQChannel struct {
	Con                           *RabbitMQConnection
	Channel                       *amqp.Channel
	NotifyPublishChan             chan amqp.Confirmation
	NotifyAckChan, NotifyNackChan chan uint64
	NotifyReturtn                 chan amqp.Return
}

func (c *RabbitMQConnection) createChannel() *RabbitMQChannel {
	if c != nil && c.Con != nil {
		channel, err := c.Con.Channel()
		if err != nil {
			panic(err)
		}
		return &RabbitMQChannel{
			Channel: channel,
			Con:     c,
		}
	}
	panic("无效的RabbitMQConnection")
}

// 获取新channel
func (c *RabbitMQConnection) GetNewChannel() *RabbitMQChannel {
	return c.createChannel()
}

// 获取发布确认的channel
func (c *RabbitMQConnection) GetNewConfirmChannel() *RabbitMQChannel {
	channel := c.createChannel()
	noWait := false
	//先申明notifyChan
	channel.NotifyPublishChan = channel.Channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	channel.NotifyAckChan, channel.NotifyNackChan = channel.Channel.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))
	channel.NotifyReturtn = channel.Channel.NotifyReturn(make(chan amqp.Return, 1))
	err := channel.Channel.Confirm(noWait)
	if err != nil {
		panic(err)
	}
	return channel
}

// 关闭channel
func (c *RabbitMQChannel) CloseChannel() {
	if c == nil {
		return
	}
	if c.Channel == nil {
		return
	}
	err := c.Channel.Close()
	if err != nil {
		panic(err)
	}
}
