package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
)

// KafkaConn - структура, определяющая соединение с Kafka-брокером.
type KafkaConn struct {
	brokers, topics []string           // список адресов брокеров, список обрабатываемых топиков
	groupID         string             // айди группы консьюмеров
	docHandler      usecase.DocHandler // обработчик документов
	numPartitions   int                //количество партиций в топике

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup
}

// NewKafkaConn создает новый экземлпяр KafkaConn.
//
// Параметры:
//   - brokers - список адресов брокеров
//   - topics - список обрабатываемых топиков
//   - groupID - айди группы консьюмеров
//   - docHandler - обработчик документов
//   - numPart - количество партиций в топике
func NewKafkaConn(brokers, topics []string, groupID string, numPart int, docHandler usecase.DocHandler) *KafkaConn {
	return &KafkaConn{
		brokers:       brokers,
		topics:        topics,
		groupID:       groupID,
		docHandler:    docHandler,
		numPartitions: numPart,
	}
}

// RunConsumer запускает консьюмеры в количестве, соответсвующей количеству партиций.
func (k *KafkaConn) RunConsumers() error {
	config := configSarama()

	consumerGroup, err := sarama.NewConsumerGroup(k.brokers, k.groupID, config)
	if err != nil {
		return fmt.Errorf("failed to start consumer group: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	k.ctx = ctx
	k.cancel = cancel

	for i := 0; i < k.numPartitions; i++ {
		k.wg.Add(1)
		go func() {
			defer k.wg.Done()
			for {
				if err := consumerGroup.Consume(k.ctx, k.topics, ConsumerGroupHandler{k.docHandler}); err != nil {
					log.Printf("error from consumer: %v", err)
				}
				if ctx.Err() != nil {
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()
	}

	log.Println("consumer is running...")
	return nil
}

// TODO
// func (k *KafkaConn) stopConsumer() {}

func configSarama() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false

	return config
}
