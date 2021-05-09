package connection

import (
	"fmt"
	"mq_demo/common"

	amqp "github.com/rabbitmq/amqp091-go"
)

var SingleConnection *RabbitMQConnection

type RabbitMQConnection struct {
	Con          *amqp.Connection
	ConCloseChan chan *amqp.Error
	ConBlockChan chan amqp.Blocking
}

func (rabbitMQConnection *RabbitMQConnection) createConnection(user string, password string, host string, port int) {
	conInfo := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)
	conn, err := amqp.Dial(conInfo)
	if err != nil {
		panic(err)
	}
	rabbitMQConnection.Con = conn
}

// 获取连接单例
func GetSingleConnection() *RabbitMQConnection {
	if SingleConnection == nil {
		SingleConnection = GetNewConnection()
		return SingleConnection
	}
	return SingleConnection
}

// 获取新连接
func GetNewConnection() *RabbitMQConnection {
	var newCon RabbitMQConnection
	user := common.RABBITMQ_USER
	password := common.RABBITMQ_PASSWOD
	host := common.RABBITMQ_HOST
	port := common.RABBITMQ_PORT
	newCon.createConnection(user, password, host, port)
	return &newCon
}

// 为连接配置监听Notify
func (c *RabbitMQConnection) ConnectionSetNotifyChannel() {
	if c != nil && c.Con != nil {
		c.ConCloseChan = c.Con.NotifyClose(make(chan *amqp.Error))
		c.ConBlockChan = c.Con.NotifyBlocked(make(chan amqp.Blocking))
	} else {
		panic("无效的RabbitMQConnection")
	}
}

// 关闭Connection
func (c *RabbitMQConnection) CloseConnection() {
	if c == nil {
		return
	}
	if c.Con == nil {
		return
	}
	c.Con.Close()
}
