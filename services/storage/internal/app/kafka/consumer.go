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

// Setup выполняется перед началом получения сообщений,
// может содержать любой функционал предподготовки консьюмера.
func (c ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup выполняется перед завершением работы консьюмера,
// может содержать любой функционал.
func (c ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim занимается получением сообщений и передачей в обработчики.
func (c ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := c.handler.HandleDocMessage(msg.Value); err != nil {
			log.Printf("failed to handle message: %v", err)
		}

		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}
