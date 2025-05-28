package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/child6yo/rago/services/storage/internal"
)

func TestDocumentTopic(t *testing.T) {
	// Настройки Kafka
	brokers := []string{"localhost:9092"}
	topic := "document-topic"

	// Создаем конфиг продюсера
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll        // Ждём ack от всех реплик
	config.Producer.Retry.Max = 5                           // Максимум попыток
	config.Producer.Return.Successes = true                 // Возвращать успехи
	config.Producer.Partitioner = sarama.NewHashPartitioner // Используем хеширование ключа для распределения

	// Создаем асинхронного продюсера
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Ошибка создания продюсера: %v", err)
	}
	defer func() {
		err := producer.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии продюсера: %v", err)
		}
	}()

	for i := 0; i <= 10; i++ {
		doc := internal.Document{
			Content: fmt.Sprintf("test%d", i),
			Metadata: internal.Metadata{
				URL: "test.com",
			},
		}
		sendEvent(producer, topic, doc)
	}

	for {
		select {
		case msg := <-producer.Successes():
			t.Logf("Сообщение отправлено в партицию %d с offset %d\n", msg.Partition, msg.Offset)
		case err := <-producer.Errors():
			t.Fatalf("Ошибка отправки сообщения: %v\n", err.Err)
		case <-time.After(2 * time.Second):
			return
		}
	}
}

func sendEvent(producer sarama.AsyncProducer, topic string, event any) {
	// Сериализуем в JSON
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		return
	}

	// Создаем сообщение
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.ByteEncoder(jsonBytes),
	}

	// Отправляем
	producer.Input() <- msg
}
