package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/child6yo/rago/services/splitter/internal"
	"github.com/child6yo/rago/services/splitter/internal/app/kafka/producer"
)

// SplitService имплементирует интерфейс Splitter.
type SplitService struct {
	wg         sync.WaitGroup
	numWorkers int // количество воркеров на втором этапе пайплайна
	producer   producer.Producer
}

// NewSplitService создает новый экземпляр SplitService.
func NewSplitService(numWorkers int, producer producer.Producer) *SplitService {
	return &SplitService{numWorkers: numWorkers, producer: producer}
}

// SplitDocuments разбивает массив документов на единичные структуры
// и асинхронно отправляет их в нижележащий сервис.
func (s *SplitService) SplitDocuments(docs []byte) error {
	uDocs, err := unmarshalDocs(docs)
	if err != nil {
		return err
	}
	s.startPipeline(uDocs)

	return nil
}

func (s *SplitService) startPipeline(docs []internal.Document) {
	doc := s.split(docs)

	for i := 0; i < s.numWorkers; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.handleSplited(doc)
		}()
	}

	s.wg.Wait()
}

func (s *SplitService) split(docs []internal.Document) <-chan internal.Document {
	out := make(chan internal.Document)
	go func() {
		defer close(out)
		for _, d := range docs {
			s.producer.SendMessage(d)
		}
	}()
	return out
}

func (s *SplitService) handleSplited(in <-chan internal.Document) {
	for val := range in {
		log.Print(val)
		s.producer.SendMessage(val)
	}
}

func unmarshalDocs(message []byte) ([]internal.Document, error) {
	if len(message) == 0 {
		return nil, errors.New("splitter: empty message")
	}

	var documents []internal.Document
	if err := json.Unmarshal(message, &documents); err != nil {
		return nil, fmt.Errorf("splitter: %w", err)
	}

	return documents, nil
}
