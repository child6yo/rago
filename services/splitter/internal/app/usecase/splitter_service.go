package usecase

import (
	"log"
	"sync"

	"github.com/child6yo/rago/services/splitter/internal"
	"github.com/child6yo/rago/services/splitter/internal/app/kafka"
	pb "github.com/child6yo/rago/services/splitter/proto"
)

// SplitService имплементирует интерфейс Splitter.
type SplitService struct {
	wg         sync.WaitGroup
	numWorkers int // количество воркеров на втором этапе пайплайна
	producer   kafka.Producer
}

// NewSplitService создает новый экземпляр SplitService.
func NewSplitService(numWorkers int, producer kafka.Producer) *SplitService {
	return &SplitService{numWorkers: numWorkers, producer: producer}
}

// SplitDocuments разбивает массив документов на единичные структуры
// и асинхронно отправляет их в нижележащий сервис.
func (s *SplitService) SplitDocuments(docs []*pb.Document) {
	s.startPipeline(docs)
}

func (s *SplitService) startPipeline(docs []*pb.Document) {
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

func (s *SplitService) split(docs []*pb.Document) <-chan internal.Document {
	out := make(chan internal.Document)
	go func() {
		defer close(out)
		for _, d := range docs {
			var doc internal.Document
			doc.Content = d.GetContent()
			doc.Metadata.URL = d.GetMetadata().GetUrl()
			out <- doc
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
