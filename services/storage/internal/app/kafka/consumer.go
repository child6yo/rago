package kafka

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/storage/internal/app/usecase"
)

// ConsumerGroupHandler имплементирует интерфейс sarama.ConsumerGroupHandler.
type ConsumerGroupHandler struct {
	handler usecase.DocHandler
}

func (c ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("Setting up consumer group session")
	return nil
}

func (c ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("Cleaning up consumer group session")
	return nil
}

func (c ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Print(msg.Partition, msg.Topic)
		if err := c.handler.HandleDocMessage(msg.Value); err != nil {
			log.Printf("Ошибка обработки сообщения: %v", err)
			continue
		}

		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}
