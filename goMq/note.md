# RABBITMQ
## 未整理
延迟队列
备份队列
惰性对列
镜像队列
TTL队列
## 事务
txSelect()：将当前channel设置成transaction模式。
txCommit()：提交事务。
txRollback()：回滚事务。

Tx puts the channel into transaction mode on the server.  All publishings and
acknowledgments following this method will be atomically committed or rolled
back for a single queue.  Call either Channel.TxCommit or Channel.TxRollback to
leave a this transaction and immediately start a new transaction.

The atomicity across multiple queues is not defined as queue declarations and
bindings are not included in the transaction.

The behavior of publishings that are delivered as mandatory or immediate while
the channel is in a transaction is not defined.

Once a channel has been put into transaction mode, it cannot be taken out of
transaction mode.  Use a different channel for non-transactional semantics.

不常用

## 各种队列

### 优先级队列
RabbitMQ一开始并没有优先级队列，而是在3.5.0版本才实现的优先级队列。

RabbitMQ的优先级队列可以让优先级高的消息先被消费者消费，优先级低的消息后被消费者消费。优先级的最大值为255，最小值为0（默认值），值越大，优先级越高，优先级越高，越先被消费者消费。

注意：优先级设置的过多，会使用更多的Erlang进程来消耗更多的CPU资源，因此，推荐优先级的值介于1和10之间
RabbitMQ不支持通过策略的方式设置队列的优先级！因此，等到流量大了，再想设置优先级，已经晚了，除非停机，但这是不可能的。

可以使用队列参数x-max-priority设置队列的最大优先级。
设置消息的优先级 Priority 

### 备份队列
### 死信队列
定义:

	一般来说，Producer 将消息投递到 Broker 或者直接到 Queue 里了，Consumer 从 Queue 取出消息进行消费，但某些时候由于特定的原因导致 Queue 中的某些消息无法被消费，这样的消息如果没有后续的处理，就变成了死信，有死信自然就有了死信队列。

应用场景:

	为了保证订单业务的消息数据不丢失，需要使用到 RabbitMQ 的死信队列机制，当消息消费发生异常时，将消息投入到死信队列中。还有比如说：用户在商城下单成功并点击支付后再指定时间未支付时自动失效。

#### 产生原因

	消息 TTL 过期
		队列统一设置过期时间 x-message-ttl
		单独设置消息的过期时间 Expiration
		如果两者都设置了过期时间，以时间短的为准。
	队列达到最大长度（队列满了，无法再添加数据到 mq 中）
		队列长度设置 x-max-legth
	消息被拒绝（basic.reject 或 basic.nack）并且 requeue=false（不再重新入队）
#### 架构
![Alt text](image.png)
### 延时队列
	通过DLX+TTL(死信队列+过期时间)
	发送信息给normal,normal不进行消费,并且配置死信队列,消息过期后,投递到死信队列,死信队列进行业务处理(完成延时)
## 公平性
在开启消费确认,RabbitMQ 提供了一种 qos （服务质量保证）功能，即在非自动确认消息的前提下，如果一定数目的消息（Channel.Qos）未被确认前，不进行消费新的消息。

	global参数:When global is true, these Qos settings apply to all existing and future consumers on all channels on the same connection. When false, the Channel.Qos settings will apply to all existing and future consumers on this channel.当次参数未ture,表示同一connection下生成的所有channel均会应用次配置,如果未false,只应用于当前channel产生的消费者
	
	To get round-robin behavior between consumers consuming from the same queue on different connections, set the prefetch count to 1, and the next available message on the server will be delivered to the next available consumer. 如果是不同的connection对同一queue的消费,配置 prefetch count=1来实现负载均衡(公平发送)

	prefetch_size 一般设置未0,表示对消息的大小不做设置

	prefetch_count 表所当前consumer在未ack之前,能够预先获取到几次消息,比如设置为2,表示该consumer能够先获取两次消息(消息的处理为异步处理,需要一定的处理时间才能够ack,能够同一时间接收大量消息)第三次获取消息时,由于前两次消息尚未ack,此时第三次消息是收不到的
## 两种消费模式
	拉模式
	channel.Get 每次获取一条
	推模式
	channel.Consume 返回一个channel,消息通过channel发送过来,可以使用range channel实现对推送的消息处理
## 可靠性

### 生产者可靠性

	1.1生产者重连
	1.2生产者确认
		1.2.1 发送到Exchange
		1.2.2 由Exchange投递给queue 
	发送给exchange的失败由DefferedConfirmWithContext判断
	而由exchange发送到queue失败由NotifyReturn监听
	NotifyReturn registers a listener for basic.return methods.  These can be sent
	from the server when a publish is undeliverable either from the mandatory or
	immediate flags.(打开mandatory 或者immediate 才能监听NotifyReturn)

### 消费者可靠性
	消费者确认
		消费者获取消息后通过ack告诉queue该消息是否已经消费
## 队列的排他性说明
exclusive 参数

Exclusive queues are only accessible by the connection that declares them and
will be deleted when the connection closes.  Channels on other connections
will receive an error when attempting  to declare, bind, consume, purge or
delete a queue with the same name.
(为true时候,即使开启了durable,一旦connection 关闭,队列就会消失,且只能被该connection访问,其他的connetion在执行 declare,bind,consume,purege,delete 该queue都会报错)
## 消费排他性说明
When exclusive is true, the server will ensure that this is the sole consumer
from this queue. When exclusive is false, the server will fairly distribute
deliveries across multiple consumers.
(Consume时候如果申明了exclusive为true,其他的consumer时不能消费该queue的)
## queue持久化说明
和队列相关的参数有两个:durable,autodelete
建议配置 durable:ture auto_delete: false


Durable and Non-Auto-Deleted queues will survive server restarts and remain when there are no remaining consumers or bindings.  Persistent publishings will
be restored in this queue on server restart.  These queues are only able to be
bound to durable exchanges.
(持久化 不自动删除, 服务重启后队列依然存在,服务没有重启前没有consumer连接也会保持,只能和持久化的exchanger绑定)

Non-Durable and Auto-Deleted queues will not be redeclared on server restart
and will be deleted by the server after a short time when the last consumer is
canceled or the last consumer's channel is closed.  Queues with this lifetime
can also be deleted normally with QueueDelete.  These durable queues can only
be bound to non-durable exchanges.
(非持久化 自动删除,服务重启后队列消失,服务没有重启前,如果没有consumer会在短时间内删除队列,只能绑定非持久化的excchanger)
Non-Durable and Non-Auto-Deleted queues will remain declared as long as the
server is running regardless of how many consumers.  This lifetime is useful
for temporary topologies that may have long delays between consumer activity.
These queues can only be bound to non-durable exchanges.
(非持久化, 不自动删除,服务重启后队列消失,服务没有重启前没有consumer连接也会保持,只能和非持久化的exchanger绑定)
Durable and Auto-Deleted queues will be restored on server restart, but without active consumers will not survive and be removed.  This Lifetime is unlikely to be useful.
(持久化,自动删除 在重启后队列会恢复,但是如果没有消费者连接队列时,队列将会被删除,该种模式不常用)
## 交换机配置说明
Each exchange belongs to one of a set of exchange kinds/types implemented by
the server. The exchange types define the functionality of the exchange - i.e.
how messages are routed through it. Once an exchange is declared, its type
cannot be changed.  The common types are "direct", "fanout", "topic" and
"headers".
(一旦交换机的类型声明后就不能更改,常见的类型有 direct fanout topic headers)

Durable and Non-Auto-Deleted exchanges will survive server restarts and remain
declared when there are no remaining bindings.  This is the best lifetime for
long-lived exchange configurations like stable routes and default exchanges.
(持久化,非自动删除 重启服务后会依然存在,在没有绑定某个queue时候会依然存在)

Non-Durable and Auto-Deleted exchanges will be deleted when there are no
remaining bindings and not restored on server restart.  This lifetime is
useful for temporary topologies that should not pollute the virtual host on
failure or after the consumers have completed.
(非持久化,自动删除,重启服务后消失,没有绑定某个queue时会删除)

Non-Durable and Non-Auto-deleted exchanges will remain as long as the server is
running including when there are no remaining bindings.  This is useful for
temporary topologies that may have long delays between bindings.
(非持久化,非自动删除,重启服务后消失,没有绑定某个queue时会保持)
Durable and Auto-Deleted exchanges will survive server restarts and will be
removed before and after server restarts when there are no remaining bindings.
These exchanges are useful for robust temporary topologies or when you require
binding durable queues to auto-deleted exchanges.
(持久化,自动删除 重启服务后出现,但是一旦没有绑定某个queue,会自动删除,适合绑定持久化,自动删除的queue)

Note: RabbitMQ declares the default exchange types like 'amq.fanout' as
durable, so queues that bind to these pre-declared exchanges must also be
durable.
(rabbitmq 提供的一些默认交换机都是持久化的,因此和这些持久化交换机绑定的queue一定也要持久化的)

Exchanges declared as `internal` do not accept publishings. Internal
exchanges are useful when you wish to implement inter-exchange topologies
that should not be exposed to users of the broker.
(internal 表示该交换机不能接受生产者发送,死信队列可以采用)
## 面试问题
### 如何保证消息的顺序性
	
	1. 一个queue,一个consumer
	2. 开启多个queue,每个queue对应一个consumer,需要顺序处理的数据依次发送给某个queue,这样这一批有序数据的消费也就有序了(也要保证消费者的顺序消费),处理这批有效数据过程中,如果出错,则对最后一条消息使用Nack(multi=ture),来对这批顺序消息都进行丢弃
	3. 事务?
### 如何保证消息不丢失
	生产者确认
	消费者确认
	高可用集群
### 如何保证消息的幂等性
	消费成功,但是消费ack丢失导致mq认为消息未被消费而重新入队,重新发送给消费者
	唯一主键


## github.com/rabbitmq/amqp091-go
QueueDeclare:

	durable: 持久化,如果是false,表明在重启后queue消失
	auto_deleted: 如果没有消费者连接,会自动删除

Pubulish:

	immediate: ture消息发送到队列发现无消费者,那么就不会存入队列,与路由键匹配的所有队列都无消费者,则将返回给生产者,false将会丢弃
	mandatory: ture交换器无法根据自身类型和路由找到一个符合的队列,那么就会返回给生产者

Delivery:
```Go
//Delivery 表示接收到的消息

type Delivery struct {
	Acknowledger Acknowledger // the channel from which this delivery arrived

	Headers Table // Application or header exchange table

	// Properties
	ContentType     string    // MIME content type
	ContentEncoding string    // MIME content encoding
	DeliveryMode    uint8     // queue implementation use - non-persistent (1) or persistent (2)
	Priority        uint8     // queue implementation use - 0 to 9
	CorrelationId   string    // application use - correlation identifier
	ReplyTo         string    // application use - address to reply to (ex: RPC)
	Expiration      string    // implementation use - message expiration spec
	MessageId       string    // application use - message identifier
	Timestamp       time.Time // application use - message timestamp
	Type            string    // application use - message type name
	UserId          string    // application use - creating user - should be authenticated user
	AppId           string    // application use - creating application id

	// Valid only with Channel.Consume
	ConsumerTag string

	// Valid only with Channel.Get
	MessageCount uint32

	DeliveryTag uint64
	Redelivered bool
	Exchange    string // basic.publish exchange
	RoutingKey  string // basic.publish routing key

	Body []byte
}

func (d Delivery) Ack(multiple bool) error //若为true，则当前消息和同一通道上所有先前未确认的消息将被确认
func (d Delivery) Reject(requeue bool) error //若为true，server会将被消费端拒绝的消息重新入列，若为false，则会丢弃消息
func (d Delivery) Nack(multiple, requeue bool) error // requeue 若为true，server会将被消费端拒绝的消息重新入列，若为false，则会丢弃消息
                                                     // multiple若为true，则当前消息和同一通道上所有先前未确认的消息将不被确认,若为false,表示当前消息不被确认

```
DeferredConfirmation

```Go
//表示发布确认
type DeferredConfirmation struct {
	DeliveryTag uint64
	// contains filtered or unexported fields
}

func (d *DeferredConfirmation) Acked() bool //返回调用该函数时候的发布是否确认,如果没有受到确认或者收到拒绝的确认消息
func (d *DeferredConfirmation) Done() <-chan struct{}
func (d *DeferredConfirmation) Wait() bool //阻塞直至发布确认,返回值true表示发布确认
func (d *DeferredConfirmation) WaitContext(ctx context.Context) (bool, error)
````

