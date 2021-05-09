# KAFKA(未看完)
## docker 安装

    docker pull bitnami/kafka
    docker pull zookeeper
    docker pull dushixiang/kafka-map

    docker network create kafka-network --driver bridge

    docker run -d --name zookeeper --network kafka-network -e ALLOW_ANONYMOUS_LOGIN=yes zookeeper:latest

    docker run -d --name kafka  --privileged=true  -e ALLOW_PLAINTEXT_LISTENER：yes --network kafka-network -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://172.18.0.3:9092 -e TZ="Asia/Shanghai" bitnami/kafka:latest

    docker run -d --name kafka-map --network kafka-network -p 9001:8080 -e DEFAULT_USERNAME=admin -e DEFAULT_PASSWORD=admin dushixiang/kafka-map:latest 

注意:

    docker对应的kafka路径:/opt/bitnami/kafka

    KAFKA_CFG_ADVERTISED_LISTENERS 指定的IP:PORT用于KAFKA-MAP连接KAFKA
## golang操作kafaka
https://www.lixueduan.com/posts/kafka/05-quick-start/
sarama/kafka  (github.com/Shopify/sarama)

kafka-go
### 异步生产者
config.Producer.Return.Errors = true    // 设定是否需要返回错误信息
config.Producer.Return.Successes = true // 设定是否需要返回成功信息
异步建议开启ERROR
同步建议二者都开启

### 生产精确一次
### 消费精确一次
通过OFFSET来实现精确消费不推荐
建议在消息中增加唯一ID,通过业务控制
## 概念

### BROKER

    一个kafka集群通常由多个broker节点组成,实现负载均衡以及容错
    broker是无状态的,通过zk来维护状态
    单个broker每秒能处理10W+数据

![alt text](image-5.png)
如图:该KAFKA中有8个BROKER,对应的一个TOPIC一共有8个PARTITION,每个PARTITION有两个REPLICATION,以第一个BROKER为例,P1是P1分区的leader,而follower存在于第二个和第八个BROKER中,P1的读写请求是第一个BROKER完成的,而第一个BROKER只负责P0和P2的数据同步

### ZK

    用来管理和协调broker,保存kafka的元数据,如topic,partition
    可以通知生产者和消费者集群的节点的添加和或者故障

### PRODUCER

### CONSUMER
![alt text](image-9.png)
分区是最小的并行单位
一个消费者可以消费多个分区
一个分区可以被多个不同的消费者组的消费者消费
一个分区不能同时被同一个消费者组的消费者消费


### CONSUMER GROUP

    CONSUMER GROUP是kafka提供的可扩展的,具有容错性机制的消费者机制
    一个CONSUMER GROUP可以包含多个CONSUMER
    一个CONSUMER GROUP有唯一的一个ID(GROUP ID)
    CONSUMER GROUP内的CONSUMER一起消费TOPIC的所有分区的数据 
    Kafka 会以 PARTITION 为单位将消息分给消费者组的各个消费者,每条消息只会被消费者组的一个消费者消费

### PARTITION
![alt text](image-4.png)
    
    kafka集群中,TOPIC被分为多个分区
    通过PARTITION对同一个TOPIC的拆分,可以使得一个TOPIC的信息存放在多个broker,提高了扩展性
### REPLICATION
![alt text](image-8.png)

    REPLICATION可以保障某个broker故障时候,数据依然可用
    一般建议主分区+副本=3(replication-factor=3)
    ISR指同步的副本集
### TOPIC 
![alt text](image-6.png)

    TOPIC是一个逻辑概念,用于发布数据,消费数据(发布到哪个TOPIC,消费哪个TOPIC)
    一旦某个消息发布到TOPIC后就不能再更新
    一个TOPIC可以包含多个PARTITION
### OFFSET
![alt text](image-7.png)

    OFFSET记录下一条将要发送给CONSUMER的消息的序号
    OFFSET默认存储在ZK中
    在一个PARTITION中,消息是顺序存储的,对应有一个增加的ID来标识,这个就是OFFSET
    OFFSET在某个PARTITION中才有意义,在PARTITION之间没有意义
### RECORD

    消息
    以KEY,VALUE形式存储
    写入消息的时候如果不指定KEY,消息会以轮询的方式写入当前TOPIC的分区
    如果是提供了不同的KEY,则KAFKA会将相同KEY的消息放入同分区

### LEADER/FOLLOWER
KAFKA会将LEADER均匀的分配到不同个的BROKER上
当LEADER崩溃后,FOLLOWER会选举出新的LEADER
### AR ISR OSR
AR 
    
    Assigned Replicas 分区的所有副本(包括LEADER本身)
ISR 
    
    所有与LEADER保持一定程度同步的副本+LEADER=ISR In Sync Replicas
OSR

    所有与LEADER相比,同步滞后过多的副本(不包括LEADER)=OSR Out Of Replicas

AR(包括LEADER)=ISR(包括LEADER)+OSR(不包括LEADER)

正常情况下 OSR为空
### CONTROLLER
KAFKA在启动后,会在所有的BROKER中选择一个CONTROLLER
CONTROLLER针对的是BROKER
TOPIC,PARTITION,REPLICATION的管理都是CONTROLLER完成的
LEADER的选举也是CONTROLLER完成的
CONTROLLER也是高可用的,一旦CONTROLLER崩溃,其他的BROKER会重新注册成新的CONTROLLER

KAFKA集群启动的时候,所有的BROKER都会尝试连接ZK,并注册自己为CONTROLLER,但是只有一个能成为CONTROLLER,其他的BROKER都会注册为CONTROLLER的监视节点

![alt text](image-21.png)
### CONSUMER GROUP REBALANCE
确保CONSUMER GROUP下的CONSUMER达成一致,分配订阅的TOPIC的每个分区的机制

触发机制:

    1.CONSUMER个数变化
    2.订阅的TOPIC发生变化
    2.订阅的TOPIC的PARTITION发生变化
缺点:
    
    1.被影响的CONSUMER GROUP下的CONSUMER共同会参与,首先会停止工作,等待REBALANCE完成
## 特性

    发布（写入）和订阅（读取）事件流，包括从其他系统持续导入/导出数据
    根据需要持久可靠地存储事件流
    在事件发生时或回顾性地处理事件流

## KAFKA作用
异步处理,系统解耦,流量削峰,日志处理
## 和其他中间件的对比
![alt text](image.png)
![alt text](image-1.png)
## KAFKA幂等性
为了实现幂等性,KAFKA引入了PID和Sequence Number
PID PRODUCER ID 每个PORDUCER初始化后,会被分配一个唯一ID,对用户可见
Sequence Number 每个生产者发送到某个TOPIC的PARTIONS的数据都对应一个从0开始自增的Number
如果一次发送的数据带有的PID和Sequence Number在PARTITION里面在PID相同的条件下,发送的Sequence Number已经小于等于保存的Sequence Number,则认为该条数据已经保存了,不会再继续保存
![alt text](image-17.png)
## KAFKA事务
KAFKA的默认事务级别是read_uncommitted(脏读)
ISOLATION_LEVEL
## 序列化

## LEADER选举
当LEARDER故障后,KAFKA会快速选举出对应的LEADER
## LEADER的负载均衡
如果某个BROKER CRASH后恢复,可能导致PARTITION的LEARER分配不均的问题,即一个BROKER存在多个PARTITION的LEADER

![alt text](image-22.png)
通过

    ./kafka-leader-election.sh --bootstrap-server SERVER_IP:PORT --topic xxx -partition=x --election-type preffered
来将某个preffered-replica设置为LEADER解决leader分配不均的问题
## KAFKA的读写流程
ZK:
    /brokers/topics/xxx_topic/partitions/xxx_partition/state
    路径下存储的了xxx_partition的leader
![alt text](image-23.png)
注意:
    
    FOLLOWER是拉取LEADER的LOG

![alt text](image-25.png)
![alt text](image-24.png)
注意:
    
    消费者消费后会提交OFFSET
## KAFKA的物理存储(未看完)
![alt text](image-26.png)
![alt text](image-27.png)

![alt text](image-29.png)
![alt text](image-28.png)
## 生产者的写入策略
轮询分区策略(默认)
随机分区策略(key为NULL,不常用)
按KEY分区策略(key不为NUL,可能出现数据倾斜,但是可以保证局部有序)
自定义分区策略

## 消费者的分区策略
range分配策略(默认)

    可以保证每个消费者消费是均衡的
    n=分区数量/消费者数量
    m=分区数量%消费者数量
    前m个消费者各消费n+1个
    剩余的消费者消费n个
    7个分区,3个消费者
    n=2
    m=1
    因此是前1个消费者各消费3个
    后面的2个消费者小为2个

round-robin策略

    将消费者组内的所有消费者以及消费者订阅的TOPIC的PARTITION按照字典顺序(topic和partition的HASHCODE进行排序),然后通过轮询的方式分配
![alt text](image-18.png)

Stricky粘性分配策略

    分区尽可能均匀
    reblance后的分配尽量和分配前尽量接近
![alt text](image-19.png)
![alt text](image-20.png)


## 消息不丢失
### 生产者ACK

    ACK 0       不要求
    ACK 1       要求LEARER写入成功
    ACK -1/ALL  要求副本同步成功
### 消费者ACK
    可以使用sarama.NewOffsetManagerFromClient提供的OffSetManager来进行OFFSET管理,
    OFFSET和GROUP_ID强相关
### 消息积压

## 消息模型
### 点对点
实现了负载均衡
![alt text](image-11.png)
### 发布订阅
![alt text](image-10.png)
### 分区与消息顺序
![alt text](image-12.png)
M3和M4是无法保证消费顺序的
![alt text](image-13.png)
可以通过设置消息的KEY来保证在不同分区存在的情况下 同一KEY的消息保存在同一个分区,这样消费的顺序性就保证了
## 消息传递语义
生产者的消息传递:

![alt text](image-14.png)
注意:

    生产者最多一次可能导致消息的丢失
    生产者至少一次可能导致消息的重复生产
消费者的消息传递:
至少一次:
![alt text](image-15.png)
至多一次:
![alt text](image-16.png)

    消费者的至少一次可能导致消息的重复消费
    消费者的至少一次是提交消费位置失败,导致下次消费的消费位置不变从而重复消费
    消费者的至多一次是消费消息失败,但是提交消费位置成功,下次消费的消费位置发生改变,消费失败的消息不会被再次消费
    
    消费者通过对OFFSET和消息消费进行绑定,当消息消费成功,此时保存当前的OFFSET,下次消费从OFFSET+1处消费即可(使用MYSQL对消费成功的数据以及OFFSET的保存放在一个事务里面)
     
## 操作
### TOPIC
创建TOPIC
    
    kafka-topics.sh --create --bootstrap-server 172.18.0.3:9092 --topic kafka-topic-test
查询TOPIC
    
    kafka-topics.sh --list --bootstrap-server 172.18.0.3:9092
TOPIC 分区扩展

    kafka-topics.sh --bootstrap-server 172.18.0.3:9092 -alter --topic kafka-topic-test --partitions 2

### PRODUCER
生产消息
    
    kafka-console-producer.sh --broker-list 172.18.0.3:9092 --topic kafka-topic-test
### CONSUMER
消费消息
    
    kafka-console-consumer.sh --bootstrap-server 172.18.0.3:9092 --topic kafka-topic-test --from-beginning
## 集群
    需要配置ZK并启动ZK
    集群中每个sesrver的broker_id都是唯一的

![alt text](image-3.png)
![alt text](image-2.png)
## 运维

### 基准测试
### 日志清理
![alt text](image-30.png)
![alt text](image-31.png)

日志是一段(segment)为单位进行清理
### 数据积压

消费超时导致消费失败(增大超时时间)
消费者OFFSET提交失败