package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/child6yo/rago/services/generator/internal"
	"github.com/child6yo/rago/services/generator/internal/app/client"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// GenerationService имплементирует интерфейс Generation.
type GenerationService struct {
	storage       client.Storage
	model         string
	ollamaAddress string
}

// NewGenerationService создает новый экземпляр GenerationService.
//
// Параметры:
//   - storage - gRPC клиент векторного хранилища
//   - model - генеративная модель (LLM)
//   - ollamaAddress - ollama URL
func NewGenerationService(storage client.Storage, model string, ollamaAddress string) *GenerationService {
	return &GenerationService{storage: storage, model: model, ollamaAddress: ollamaAddress}
}

// Generate генерирует ответ за запрос. Самостоятельно идет в сервис хранения данных.
// Ответ генерирует порционно. Возвращает канал, в который будет стримить поток ответа.
// Возвращает ошибку в случае неполадок.
func (gs *GenerationService) Generate(ctx context.Context, query string) (<-chan string, error) {
	docs, err := gs.storage.Search(context.Background(), query, 1)
	if err != nil {
		return nil, err
	}

	contextJSON, err := prepareContext(docs)
	if err != nil {
		return nil, err
	}

	out := make(chan string)

	go func() {
		gs.initGenerating(ctx, out, query, contextJSON)
	}()

	return out, nil
}

func (gs *GenerationService) initGenerating(ctx context.Context, out chan<- string, query string, contextJSON string) {
	llm, err := ollama.New(ollama.WithModel(gs.model), ollama.WithServerURL(gs.ollamaAddress))
	if err != nil {
		log.Printf("generation failed: %v", err)
	}
	_, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		fmt.Sprintf(defaultPrompt, query, contextJSON),
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			streamingFunc(ctx, out, chunk)
			return nil
		}),
	)
	if err != nil {
		log.Printf("generation failed: %v", err)
	}

	defer close(out)
}

func streamingFunc(ctx context.Context, out chan<- string, chunk []byte) {
	if ctx.Err() != nil {
		return
	}
	out <- string(chunk)
}

func prepareContext(contextDocs []internal.Document) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("generation error: failed to prepare context: %v", err)
		}
	}()

	res, err := json.Marshal(contextDocs)

	return string(res), err
}
