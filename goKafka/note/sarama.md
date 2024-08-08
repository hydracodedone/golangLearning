# sarama 代码解析
## 生产者代码解析
https://www.lixueduan.com/posts/kafka/06-sarama-producer/
## 配置文件
```GO
const (
	// NoResponse doesn't send any response, the TCP ACK is all you get.
	NoResponse RequiredAcks = 0
	// WaitForLocal waits for only the local commit to succeed before responding.
	WaitForLocal RequiredAcks = 1
	// WaitForAll waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.insync.replicas` configuration key.
	WaitForAll RequiredAcks = -1
)


type Config struct {
	// Admin is the namespace for ClusterAdmin properties used by the administrative Kafka client.
	Admin struct {
		Retry struct {
			// The total number of times to retry sending (retriable) admin requests (default 5).
			// Similar to the `retries` setting of the JVM AdminClientConfig.
			Max int
			// Backoff time between retries of a failed request (default 100ms)
			Backoff time.Duration
		}
		// The maximum duration the administrative Kafka client will wait for ClusterAdmin operations,
		// including topics, brokers, configurations and ACLs (defaults to 3 seconds).
		Timeout time.Duration
	}

	// Net is the namespace for network-level properties used by the Broker, and
	// shared by the Client/Producer/Consumer.
	Net struct {
		// How many outstanding requests a connection is allowed to have before
		// sending on it blocks (default 5).
		// Throughput can improve but message ordering is not guaranteed if Producer.Idempotent is disabled, see:
		// https://kafka.apache.org/protocol#protocol_network
		// https://kafka.apache.org/28/documentation.html#producerconfigs_max.in.flight.requests.per.connection
		//参数代表了允许没有收到acks而可以同时发送的最大batch数
		MaxOpenRequests int

		// All three of the below configurations are similar to the
		// `socket.timeout.ms` setting in JVM kafka. All of them default
		// to 30 seconds.
		DialTimeout  time.Duration // How long to wait for the initial connection.
		ReadTimeout  time.Duration // How long to wait for a response.
		WriteTimeout time.Duration // How long to wait for a transmit.

		// ResolveCanonicalBootstrapServers turns each bootstrap broker address
		// into a set of IPs, then does a reverse lookup on each one to get its
		// canonical hostname. This list of hostnames then replaces the
		// original address list. Similar to the `client.dns.lookup` option in
		// the JVM client, this is especially useful with GSSAPI, where it
		// allows providing an alias record instead of individual broker
		// hostnames. Defaults to false.
		ResolveCanonicalBootstrapServers bool

		TLS struct {
			// Whether or not to use TLS when connecting to the broker
			// (defaults to false).
			Enable bool
			// The TLS configuration to use for secure connections if
			// enabled (defaults to nil).
			Config *tls.Config
		}

		// SASL based authentication with broker. While there are multiple SASL authentication methods
		// the current implementation is limited to plaintext (SASL/PLAIN) authentication
		SASL struct {
			// Whether or not to use SASL authentication when connecting to the broker
			// (defaults to false).
			Enable bool
			// SASLMechanism is the name of the enabled SASL mechanism.
			// Possible values: OAUTHBEARER, PLAIN (defaults to PLAIN).
			Mechanism SASLMechanism
			// Version is the SASL Protocol Version to use
			// Kafka > 1.x should use V1, except on Azure EventHub which use V0
			Version int16
			// Whether or not to send the Kafka SASL handshake first if enabled
			// (defaults to true). You should only set this to false if you're using
			// a non-Kafka SASL proxy.
			Handshake bool
			// AuthIdentity is an (optional) authorization identity (authzid) to
			// use for SASL/PLAIN authentication (if different from User) when
			// an authenticated user is permitted to act as the presented
			// alternative user. See RFC4616 for details.
			AuthIdentity string
			// User is the authentication identity (authcid) to present for
			// SASL/PLAIN or SASL/SCRAM authentication
			User string
			// Password for SASL/PLAIN authentication
			Password string
			// authz id used for SASL/SCRAM authentication
			SCRAMAuthzID string
			// SCRAMClientGeneratorFunc is a generator of a user provided implementation of a SCRAM
			// client used to perform the SCRAM exchange with the server.
			SCRAMClientGeneratorFunc func() SCRAMClient
			// TokenProvider is a user-defined callback for generating
			// access tokens for SASL/OAUTHBEARER auth. See the
			// AccessTokenProvider interface docs for proper implementation
			// guidelines.
			TokenProvider AccessTokenProvider

			GSSAPI GSSAPIConfig
		}

		// KeepAlive specifies the keep-alive period for an active network connection (defaults to 0).
		// If zero or positive, keep-alives are enabled.
		// If negative, keep-alives are disabled.
		KeepAlive time.Duration

		// LocalAddr is the local address to use when dialing an
		// address. The address must be of a compatible type for the
		// network being dialed.
		// If nil, a local address is automatically chosen.
		LocalAddr net.Addr

		Proxy struct {
			// Whether or not to use proxy when connecting to the broker
			// (defaults to false).
			Enable bool
			// The proxy dialer to use enabled (defaults to nil).
			Dialer proxy.Dialer
		}
	}

	// Metadata is the namespace for metadata management properties used by the
	// Client, and shared by the Producer/Consumer.
	Metadata struct {
		Retry struct {
			// The total number of times to retry a metadata request when the
			// cluster is in the middle of a leader election (default 3).
			Max int
			// How long to wait for leader election to occur before retrying
			// (default 250ms). Similar to the JVM's `retry.backoff.ms`.
			Backoff time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc func(retries, maxRetries int) time.Duration
		}
		// How frequently to refresh the cluster metadata in the background.
		// Defaults to 10 minutes. Set to 0 to disable. Similar to
		// `topic.metadata.refresh.interval.ms` in the JVM version.
		RefreshFrequency time.Duration

		// Whether to maintain a full set of metadata for all topics, or just
		// the minimal set that has been necessary so far. The full set is simpler
		// and usually more convenient, but can take up a substantial amount of
		// memory if you have many topics and partitions. Defaults to true.
		Full bool

		// How long to wait for a successful metadata response.
		// Disabled by default which means a metadata request against an unreachable
		// cluster (all brokers are unreachable or unresponsive) can take up to
		// `Net.[Dial|Read]Timeout * BrokerCount * (Metadata.Retry.Max + 1) + Metadata.Retry.Backoff * Metadata.Retry.Max`
		// to fail.
		Timeout time.Duration

		// Whether to allow auto-create topics in metadata refresh. If set to true,
		// the broker may auto-create topics that we requested which do not already exist,
		// if it is configured to do so (`auto.create.topics.enable` is true). Defaults to true.
		AllowAutoTopicCreation bool
	}

	// Producer is the namespace for configuration related to producing messages,
	// used by the Producer.
	Producer struct {
		// The maximum permitted size of a message (defaults to 1000000). Should be
		// set equal to or smaller than the broker's `message.max.bytes`.
		//数影响了一条消息的最大字节数，默认是1000000。但是注意，这个参数必须要小于broker中的 message.max.bytes
		MaxMessageBytes int
		// The level of acknowledgement reliability needed from the broker (defaults
		// to WaitForLocal). Equivalent to the `request.required.acks` setting of the
		// JVM producer.
		RequiredAcks RequiredAcks
		// The maximum duration the broker will wait the receipt of the number of
		// RequiredAcks (defaults to 10 seconds). This is only relevant when
		// RequiredAcks is set to WaitForAll or a number > 1. Only supports
		// millisecond resolution, nanoseconds will be truncated. Equivalent to
		// the JVM producer's `request.timeout.ms` setting.
		Timeout time.Duration
		// The type of compression to use on messages (defaults to no compression).
		// Similar to `compression.codec` setting of the JVM producer.
		Compression CompressionCodec
		// The level of compression to use on messages. The meaning depends
		// on the actual compression type used and defaults to default compression
		// level for the codec.
		CompressionLevel int
		// Generates partitioners for choosing the partition to send messages to
		// (defaults to hashing the message key). Similar to the `partitioner.class`
		// setting for the JVM producer.
		Partitioner PartitionerConstructor
		// If enabled, the producer will ensure that exactly one copy of each message is
		// written.
        //当将`Idempotent`参数设置为`true`时，生产者将会启用幂等性。这意味着生产者会在发送消息时生成一个唯一的ID，并在重试时使用相同的ID来确保消息不会被重复发送。此外，启用幂等性还会自动设置`RequiredAcks`参数为`-1`，以确保消息只有在所有ISR中的分区副本都成功写入后才被确认。
        //当将`Idempotent`参数设置为`false`时，生产者将不启用幂等性。这意味着生产者发送的消息可能会出现重复发送的情况，因为Kafka无法保证消息的幂等性。
        //启用幂等性可以确保消息在发送过程中不会被重复处理，从而提高消息传递的可靠性。然而，需要注意的是，启用幂等性会增加一些额外的开销，因此在一些特定的场景下可能需要权衡考虑。
    Idempotent bool
		Idempotent bool
		// Transaction specify
		Transaction struct {
			// Used in transactions to identify an instance of a producer through restarts
			ID string
			// Amount of time a transaction can remain unresolved (neither committed nor aborted)
			// default is 1 min
			Timeout time.Duration

			Retry struct {
				// The total number of times to retry sending a message (default 50).
				// Similar to the `message.send.max.retries` setting of the JVM producer.
				Max int
				// How long to wait for the cluster to settle between retries
				// (default 10ms). Similar to the `retry.backoff.ms` setting of the
				// JVM producer.
				Backoff time.Duration
				// Called to compute backoff time dynamically. Useful for implementing
				// more sophisticated backoff strategies. This takes precedence over
				// `Backoff` if set.
				BackoffFunc func(retries, maxRetries int) time.Duration
			}
		}

		// Return specifies what channels will be populated. If they are set to true,
		// you must read from the respective channels to prevent deadlock. If,
		// however, this config is used to create a `SyncProducer`, both must be set
		// to true and you shall not read from the channels since the producer does
		// this internally.
		Return struct {
			// If enabled, successfully delivered messages will be returned on the
			// Successes channel (default disabled).
			Successes bool

			// If enabled, messages that failed to deliver will be returned on the
			// Errors channel, including error (default enabled).
			Errors bool
		}

		// The following config options control how often messages are batched up and
		// sent to the broker. By default, messages are sent as fast as possible, and
		// all messages received while the current batch is in-flight are placed
		// into the subsequent batch.
		//用于设置将消息打包发送，简单来讲就是每次发送消息到broker的时候，不是生产一条消息就发送一条消息，而是等消息累积到一定的程度了，再打包发送。所以里面含有两个参数。一个是多少条消息触发打包发送，一个是累计的消息大小到了多少，然后发送。
		Flush struct {
			// The best-effort number of bytes needed to trigger a flush. Use the
			// global sarama.MaxRequestSize to set a hard upper limit.
			Bytes int
			// The best-effort number of messages needed to trigger a flush. Use
			// `MaxMessages` to set a hard upper limit.
			Messages int
			// The best-effort frequency of flushes. Equivalent to
			// `queue.buffering.max.ms` setting of JVM producer.
			Frequency time.Duration
			// The maximum number of messages the producer will send in a single
			// broker request. Defaults to 0 for unlimited. Similar to
			// `queue.buffering.max.messages` in the JVM producer.
			MaxMessages int
		}

		Retry struct {
			// The total number of times to retry sending a message (default 3).
			// Similar to the `message.send.max.retries` setting of the JVM producer.
			Max int
			// How long to wait for the cluster to settle between retries
			// (default 100ms). Similar to the `retry.backoff.ms` setting of the
			// JVM producer.
			Backoff time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc func(retries, maxRetries int) time.Duration
		}

		// Interceptors to be called when the producer dispatcher reads the
		// message for the first time. Interceptors allows to intercept and
		// possible mutate the message before they are published to Kafka
		// cluster. *ProducerMessage modified by the first interceptor's
		// OnSend() is passed to the second interceptor OnSend(), and so on in
		// the interceptor chain.
		Interceptors []ProducerInterceptor
	}

	// Consumer is the namespace for configuration related to consuming messages,
	// used by the Consumer.
	Consumer struct {
		// Group is the namespace for configuring consumer group.
		Group struct {
			Session struct {
				// The timeout used to detect consumer failures when using Kafka's group management facility.
				// The consumer sends periodic heartbeats to indicate its liveness to the broker.
				// If no heartbeats are received by the broker before the expiration of this session timeout,
				// then the broker will remove this consumer from the group and initiate a rebalance.
				// Note that the value must be in the allowable range as configured in the broker configuration
				// by `group.min.session.timeout.ms` and `group.max.session.timeout.ms` (default 10s)
				Timeout time.Duration
			}
			Heartbeat struct {
				// The expected time between heartbeats to the consumer coordinator when using Kafka's group
				// management facilities. Heartbeats are used to ensure that the consumer's session stays active and
				// to facilitate rebalancing when new consumers join or leave the group.
				// The value must be set lower than Consumer.Group.Session.Timeout, but typically should be set no
				// higher than 1/3 of that value.
				// It can be adjusted even lower to control the expected time for normal rebalances (default 3s)
				Interval time.Duration
			}
			Rebalance struct {
				// Strategy for allocating topic partitions to members.
				// Deprecated: Strategy exists for historical compatibility
				// and should not be used. Please use GroupStrategies.
				Strategy BalanceStrategy

				// GroupStrategies is the priority-ordered list of client-side consumer group
				// balancing strategies that will be offered to the coordinator. The first
				// strategy that all group members support will be chosen by the leader.
				// default: [ NewBalanceStrategyRange() ]
				GroupStrategies []BalanceStrategy

				// The maximum allowed time for each worker to join the group once a rebalance has begun.
				// This is basically a limit on the amount of time needed for all tasks to flush any pending
				// data and commit offsets. If the timeout is exceeded, then the worker will be removed from
				// the group, which will cause offset commit failures (default 60s).
				Timeout time.Duration

				Retry struct {
					// When a new consumer joins a consumer group the set of consumers attempt to "rebalance"
					// the load to assign partitions to each consumer. If the set of consumers changes while
					// this assignment is taking place the rebalance will fail and retry. This setting controls
					// the maximum number of attempts before giving up (default 4).
					Max int
					// Backoff time between retries during rebalance (default 2s)
					Backoff time.Duration
				}
			}
			Member struct {
				// Custom metadata to include when joining the group. The user data for all joined members
				// can be retrieved by sending a DescribeGroupRequest to the broker that is the
				// coordinator for the group.
				UserData []byte
			}

			// support KIP-345
			InstanceId string

			// If true, consumer offsets will be automatically reset to configured Initial value
			// if the fetched consumer offset is out of range of available offsets. Out of range
			// can happen if the data has been deleted from the server, or during situations of
			// under-replication where a replica does not have all the data yet. It can be
			// dangerous to reset the offset automatically, particularly in the latter case. Defaults
			// to true to maintain existing behavior.
			ResetInvalidOffsets bool
		}

		Retry struct {
			// How long to wait after a failing to read from a partition before
			// trying again (default 2s).
			Backoff time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc func(retries int) time.Duration
		}

		// Fetch is the namespace for controlling how many bytes are retrieved by any
		// given request.
		Fetch struct {
			// The minimum number of message bytes to fetch in a request - the broker
			// will wait until at least this many are available. The default is 1,
			// as 0 causes the consumer to spin when no messages are available.
			// Equivalent to the JVM's `fetch.min.bytes`.
			Min int32
			// The default number of message bytes to fetch from the broker in each
			// request (default 1MB). This should be larger than the majority of
			// your messages, or else the consumer will spend a lot of time
			// negotiating sizes and not actually consuming. Similar to the JVM's
			// `fetch.message.max.bytes`.
			Default int32
			// The maximum number of message bytes to fetch from the broker in a
			// single request. Messages larger than this will return
			// ErrMessageTooLarge and will not be consumable, so you must be sure
			// this is at least as large as your largest message. Defaults to 0
			// (no limit). Similar to the JVM's `fetch.message.max.bytes`. The
			// global `sarama.MaxResponseSize` still applies.
			Max int32
		}
		// The maximum amount of time the broker will wait for Consumer.Fetch.Min
		// bytes to become available before it returns fewer than that anyways. The
		// default is 250ms, since 0 causes the consumer to spin when no events are
		// available. 100-500ms is a reasonable range for most cases. Kafka only
		// supports precision up to milliseconds; nanoseconds will be truncated.
		// Equivalent to the JVM's `fetch.wait.max.ms`.
		MaxWaitTime time.Duration

		// The maximum amount of time the consumer expects a message takes to
		// process for the user. If writing to the Messages channel takes longer
		// than this, that partition will stop fetching more messages until it
		// can proceed again.
		// Note that, since the Messages channel is buffered, the actual grace time is
		// (MaxProcessingTime * ChannelBufferSize). Defaults to 100ms.
		// If a message is not written to the Messages channel between two ticks
		// of the expiryTicker then a timeout is detected.
		// Using a ticker instead of a timer to detect timeouts should typically
		// result in many fewer calls to Timer functions which may result in a
		// significant performance improvement if many messages are being sent
		// and timeouts are infrequent.
		// The disadvantage of using a ticker instead of a timer is that
		// timeouts will be less accurate. That is, the effective timeout could
		// be between `MaxProcessingTime` and `2 * MaxProcessingTime`. For
		// example, if `MaxProcessingTime` is 100ms then a delay of 180ms
		// between two messages being sent may not be recognized as a timeout.
		MaxProcessingTime time.Duration

		// Return specifies what channels will be populated. If they are set to true,
		// you must read from them to prevent deadlock.
		Return struct {
			// If enabled, any errors that occurred while consuming are returned on
			// the Errors channel (default disabled).
			Errors bool
		}

		// Offsets specifies configuration for how and when to commit consumed
		// offsets. This currently requires the manual use of an OffsetManager
		// but will eventually be automated.
		Offsets struct {
			// Deprecated: CommitInterval exists for historical compatibility
			// and should not be used. Please use Consumer.Offsets.AutoCommit
			CommitInterval time.Duration

			// AutoCommit specifies configuration for commit messages automatically.
			AutoCommit struct {
				// Whether or not to auto-commit updated offsets back to the broker.
				// (default enabled).
				Enable bool

				// How frequently to commit updated offsets. Ineffective unless
				// auto-commit is enabled (default 1s)
				Interval time.Duration
			}

			// The initial offset to use if no offset was previously committed.
			// Should be OffsetNewest or OffsetOldest. Defaults to OffsetNewest.
			Initial int64

			// The retention duration for committed offsets. If zero, disabled
			// (in which case the `offsets.retention.minutes` option on the
			// broker will be used).  Kafka only supports precision up to
			// milliseconds; nanoseconds will be truncated. Requires Kafka
			// broker version 0.9.0 or later.
			// (default is 0: disabled).
			Retention time.Duration

			Retry struct {
				// The total number of times to retry failing commit
				// requests during OffsetManager shutdown (default 3).
				Max int
			}
		}

		// IsolationLevel support 2 mode:
		// 	- use `ReadUncommitted` (default) to consume and return all messages in message channel
		//	- use `ReadCommitted` to hide messages that are part of an aborted transaction
		IsolationLevel IsolationLevel

		// Interceptors to be called just before the record is sent to the
		// messages channel. Interceptors allows to intercept and possible
		// mutate the message before they are returned to the client.
		// *ConsumerMessage modified by the first interceptor's OnConsume() is
		// passed to the second interceptor OnConsume(), and so on in the
		// interceptor chain.
		Interceptors []ConsumerInterceptor
	}

	// A user-provided string sent with every request to the brokers for logging,
	// debugging, and auditing purposes. Defaults to "sarama", but you should
	// probably set it to something specific to your application.
	ClientID string
	// A rack identifier for this client. This can be any string value which
	// indicates where this client is physically located.
	// It corresponds with the broker config 'broker.rack'
	RackID string
	// The number of events to buffer in internal and external channels. This
	// permits the producer and consumer to continue processing some messages
	// in the background while user code is working, greatly improving throughput.
	// Defaults to 256.
	ChannelBufferSize int
	// ApiVersionsRequest determines whether Sarama should send an
	// ApiVersionsRequest message to each broker as part of its initial
	// connection. This defaults to `true` to match the official Java client
	// and most 3rdparty ones.
	ApiVersionsRequest bool
	// The version of Kafka that Sarama will assume it is running against.
	// Defaults to the oldest supported stable version. Since Kafka provides
	// backwards-compatibility, setting it to a version older than you have
	// will not break anything, although it may prevent you from using the
	// latest features. Setting it to a version greater than you are actually
	// running may lead to random breakage.
	Version KafkaVersion
	// The registry to define metrics into.
	// Defaults to a local registry.
	// If you want to disable metrics gathering, set "metrics.UseNilMetrics" to "true"
	// prior to starting Sarama.
	// See Examples on how to use the metrics registry
	MetricRegistry metrics.Registry
}

```
## panic 处理
```GO
func withRecover(fn func()) {
	defer func() {
		handler := PanicHandler
		if handler != nil {
			if err := recover(); err != nil {
				handler(err)
			}
		}
	}()

	fn()
}
```
##