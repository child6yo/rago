package consumer

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
)

// Connection - структура, определяющая соединение с Kafka-брокером.
type Connection struct {
	brokers, topics []string         // список адресов брокеров, список обрабатываемых топиков
	groupID         string           // айди группы консьюмеров
	handler         usecase.Splitter // обработчик документов
	numPartitions   int              //количество партиций в топике

	ctx    *context.Context
	cancel *context.CancelFunc

	wg sync.WaitGroup
}

// NewConnection создает новый экземлпяр Connection.
//
// Параметры:
//   - brokers - список адресов брокеров
//   - topics - список обрабатываемых топиков
//   - groupID - айди группы консьюмеров
//   - docHandler - обработчик документов
//   - numPart - количество партиций в топике
func NewConnection(brokers, topics []string, groupID string, numPart int, handler usecase.Splitter) *Connection {
	return &Connection{
		brokers:       brokers,
		topics:        topics,
		groupID:       groupID,
		handler:       handler,
		numPartitions: numPart,
	}
}

// RunConsumers запускает консьюмеры в количестве, соответсвующем количеству партиций.
func (c *Connection) RunConsumers() error {
	config := configSarama()

	consumerGroup, err := sarama.NewConsumerGroup(c.brokers, c.groupID, config)
	if err != nil {
		return fmt.Errorf("splitter run consumers: failed to start consumer group: %w", err)
	}

	go func() {
		for err := range consumerGroup.Errors() {
			log.Printf("splitter consumer error: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = &ctx
	c.cancel = &cancel

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, c.topics, ConsumerGroupHandler{c.handler}); err != nil {
				log.Printf("splitter run consumers: error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	log.Println("splitter: consumer group is running...")
	return nil
}

// StopConsumers останавливает группу консьюмеров отменой контекста.
// Дожидается завершения всех горутин.
func (c *Connection) StopConsumers() {
	cnl := *c.cancel
	cnl()

	c.wg.Wait()
}
