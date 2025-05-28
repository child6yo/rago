package app

import (
	"log"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct {
	*Application
}

func (h ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("Setting up consumer group session")
	return nil
}

func (h ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("Cleaning up consumer group session")
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("received message: key=%s, value=%s, partition=%d, offset=%d\n",
			message.Key, message.Value, message.Partition, message.Offset)

		err := h.handleDocMessage(message)
		if err != nil {
			// TODO
			log.Println("later")
		}
		session.MarkMessage(message, "")
	}
	return nil
}
