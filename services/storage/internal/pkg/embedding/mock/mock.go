package mock

import "context"

// Embedder mock
type Embedder struct {
	GenerateEmbeddingsFunc func(ctx context.Context, input string) ([]float32, error)
}

// GenerateEmbeddings mock
func (e *Embedder) GenerateEmbeddings(ctx context.Context, input string) ([]float32, error) {
	if e.GenerateEmbeddingsFunc != nil {
		return e.GenerateEmbeddingsFunc(ctx, input)
	}
	return []float32{}, nil
}
