package embedding

import "context"

// Embedder определяет интерфейс взаимодействия с эмбеддинговой моделью.
type Embedder interface {
	// GenerateEmbeddings генерирует эмбеддинговое представление передаваемых текстовых данных.
	GenerateEmbeddings(ctx context.Context, input string) ([]float32, error)
}