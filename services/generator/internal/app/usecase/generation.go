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

type GenerationService struct {
	storage client.Storage
}

func NewGenerationService(storage client.Storage) *GenerationService {
	return &GenerationService{storage: storage}
}

// Generate генерирует ответ за запрос. Самостоятельно идет в сервис хранения данных.
// Ответ генерирует порционно. Возвращает канал, в который будет стримить поток ответа.
// Возвращает ошибку в случае неполадок.
func (gs *GenerationService) Generate(ctx context.Context, query string) (<-chan string, error) {
	docs, err := gs.storage.Search(context.Background(), query, 1)
	if err != nil {
		return nil, err
	}

	contextJson, err := prepareContext(docs)
	if err != nil {
		return nil, err
	}

	out := make(chan string)

	go func() {
		gs.initGenerating(ctx, out, query, contextJson)
	}()

	return out, nil
}

func (gs *GenerationService) initGenerating(ctx context.Context, out chan<- string, query string, contextJson string) {
	llm, err := ollama.New(ollama.WithModel("gemma3:1b"), ollama.WithServerURL("http://localhost:11434"))
	if err != nil {
		log.Println(err)
	}
	_, err = llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		fmt.Sprintf(defaultPrompt, query, contextJson),
		llms.WithTemperature(0.2),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			streamingFunc(out, ctx, chunk)
			return nil
		}),
	)
	if err != nil {
		log.Println(err)
	}

	defer close(out)
}

func streamingFunc(out chan<- string, ctx context.Context, chunk []byte) {
	if ctx.Err() != nil {
		return
	}
	out <- string(chunk)
}

func prepareContext(contextDocs []internal.Document) (string, error) {
	res, err := json.Marshal(contextDocs)

	return string(res), err
}
