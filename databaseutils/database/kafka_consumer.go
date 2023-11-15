package database

import (
	"context"
	"github.com/IBM/sarama"
	"sync"
)

type (
	KafkaConsumer struct {
		client        sarama.ConsumerGroup // 客户端链接
		stopped       chan struct{}        // 停止通道
		context       context.Context      // 上下文
		cancel        context.CancelFunc
		wg            sync.WaitGroup
		initCfg       *KafkaConsumerInitCfg
		consumerClaim ConsumerClaim
	}

	KafkaConsumerInitCfg struct {
		KafkaName         string                       // kafka 配置名称
		GroupId           string                       // 消费组名称
		Topics            string                       // 需要订阅的主题
		MessageBufferSize int                          // message的缓冲区
		Context           context.Context              // 上下文
		ResetOffsetCfg    map[string][]PartitionOffset // 自定义topic主题下的partition偏移量
		Earliest          bool
	}

	PartitionOffset struct {
		Topic     string
		Partition int32 // 分区
		Offset    int64 // 偏移量
	}

	ConsumerClaim struct {
		msgs           chan *sarama.ConsumerMessage
		resetOffsetCfg map[string][]PartitionOffset
		ready          chan bool
	}
)
