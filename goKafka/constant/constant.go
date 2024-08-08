package constant

import "errors"

var KAFKA_ADDRESS = "localhost:9092"

var TOPIC = "GO_KAFKA_TOPIC"
var KEY = "GO_KAFKA_MESSAGE_KEY"
var PARTITION_0 int32 = 0
var PARTITION_1 int32 = 1
var CONSUMER_GROUP = "CONSUMER_GROUP_FOR_GO_KAFKA_TOPIC"

var ErrKafakaProducerNotInitialized = errors.New("kafka producer not initialized")
var ErrKafakaConsusmerNotInitialized = errors.New("kafka consumer not initialized")

var ErrKafakaProducerHasClosed = errors.New("kafka producer has closed")
var ErrKafakaProducerSendMessageTimeout = errors.New("kafka producer send message timeout")
