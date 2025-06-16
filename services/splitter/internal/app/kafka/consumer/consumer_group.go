package consumer

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/splitter/internal/app/usecase"
)

// ConsumerGroupHandler имплементирует интерфейс sarama.ConsumerGroupHandler.
type ConsumerGroupHandler struct {
	handler usecase.Splitter
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
		log.Printf("splitter consumer: message recieved from partition %d, offset %d", msg.Partition, msg.Offset)
		err := c.handler.SplitDocuments(msg.Value)
		if err != nil {
			return err
		}

		log.Printf("splitter consumer: message successfully handled, offset %d", msg.Offset)
		session.MarkMessage(msg, "")
	}
	return nil
}
