package producer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/splitter/internal"
)

// Producer описывает интерфейс продюсера, способного отправлять
// сообщения в брокер.
type Producer interface {
	// SendMessage отправляет сообщение в определенный топик брокера.
	SendMessage(event internal.Document)
}

// KafkaProducer имплементирует интерфейс Producer.
type KafkaProducer struct {
	producer sarama.AsyncProducer
	brokers  []string
	docTopic string
}

// NewKafkaProducer создает новый экземпляр KafkaProducer.
//
// Параметры:
//   - brokers - слайс с адресами брокеров
//   - docTopic - название топика
func NewKafkaProducer(brokers []string, docTopic string) *KafkaProducer {
	return &KafkaProducer{
		brokers:  brokers,
		docTopic: docTopic,
	}
}

// StartProducer запускает работу продюсера.
func (p *KafkaProducer) StartProducer() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewAsyncProducer(p.brokers, config)
	if err != nil {
		return fmt.Errorf("failed to start producer: %v", err)
	}

	p.producer = producer

	// Горутина для обработки успехов
	go func() {
		for msg := range producer.Successes() {
			log.Printf("document successfully sended: partition=%d, offset=%d\n", msg.Partition, msg.Offset)
		}
	}()

	// Горутина для обработки ошибок
	go func() {
		for err := range producer.Errors() {
			log.Printf("failed to send document: %v\n", err.Err)
		}
	}()

	return nil
}

// StopProducer останавливает отправку сообщений продюсером.
func (p *KafkaProducer) StopProducer() {
	if err := p.producer.Close(); err != nil {
		log.Printf("failed to gracefully stop producer: %v", err)
	}
}

// SendMessage отправляет сообщение в брокер.
func (p *KafkaProducer) SendMessage(event internal.Document) {
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("consumer send message: failed do marshal JSON: %v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: p.docTopic,
		Value: sarama.ByteEncoder(jsonBytes),
	}

	p.producer.Input() <- msg
}
