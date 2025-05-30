package kafka

import (
	"context"
	"encoding/json"
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
	ctx      context.Context
	cancel   context.CancelFunc
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

// SendMessage отправляет сообщение в определенный топик брокера.
func (p *KafkaProducer) StartProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewAsyncProducer(p.brokers, config)
	if err != nil {
		log.Fatalf("Ошибка создания продюсера: %v", err)
	}

	p.producer = producer

	// defer func() {
	// 	if err := producer.Close(); err != nil {
	// 		log.Printf("Ошибка при закрытии продюсера: %v", err)
	// 	}
	// }()

	// Горутина для обработки успехов
	go func() {
		for msg := range producer.Successes() {
			log.Printf("Сообщение успешно отправлено: partition=%d, offset=%d\n", msg.Partition, msg.Offset)
		}
	}()

	// Горутина для обработки ошибок
	go func() {
		for err := range producer.Errors() {
			log.Printf("Ошибка при отправке сообщения: %v\n", err.Err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	p.ctx = ctx
	p.cancel = cancel

	// log.Println("Отправка сообщений... Нажмите Ctrl+C для выхода")
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
	// log.Println("Завершаем работу...")
	// cancel()
	// wg.Wait()
}

func (p *KafkaProducer) SendMessage(event internal.Document) {
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.docTopic,
		Value: sarama.ByteEncoder(jsonBytes),
	}

	p.producer.Input() <- msg
}
