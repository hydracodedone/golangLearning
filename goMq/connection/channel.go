package connection

import (
    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQChannel struct {
    // Con 指向关联的 RabbitMQConnection
    Con *RabbitMQConnection
    // Channel 实际的 AMQP Channel
    Channel *amqp.Channel
    // NotifyPublishChan 发布确认的通知通道
    NotifyPublishChan chan amqp.Confirmation
    // NotifyAckChan/NotifyNackChan 分别用于接收 ACK 和 NACK 的 delivery tag
    NotifyAckChan, NotifyNackChan chan uint64
    // NotifyReturn return 消息的通知通道
    NotifyReturn chan amqp.Return
}

// createChannel 创建一个新的 AMQP Channel 并封装为 RabbitMQChannel
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
    channel.NotifyReturn = channel.Channel.NotifyReturn(make(chan amqp.Return, 1))
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
