package app

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

// TODO
func runConsumer(brokers []string, groupID string, topics []string, app *Application) error {
	config := configSarama()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return fmt.Errorf("failed to start consumer group: %w", err)
	}

	ctx := context.Background()
	app.kfkWG = &sync.WaitGroup{}

	app.kfkWG.Add(app.NumWorkers)
	for i := 0; i <= app.NumWorkers; i++ {
		go func() {
			defer app.kfkWG.Done()
			for {
				if err := consumerGroup.Consume(ctx, topics, ConsumerGroupHandler{app}); err != nil {
					log.Printf("error from consumer: %v", err)
				}

				if ctx.Err() != nil {
					return
				}
			}
		}()
	}

	log.Println("consumer is running...")

	return nil
}

// TODO
func stopConsumer() {}

// comment
func configSarama() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false

	return config
}
