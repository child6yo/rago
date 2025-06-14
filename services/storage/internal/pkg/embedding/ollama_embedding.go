package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const embeddingsURL = "/api/embeddings"

// OllamaResponse определяет ответ ollama.
type OllamaResponse struct {
	Embedding []float32 `json:"embedding"`	
}

// Ollama имплементирует интерфейс Embedder.
type OllamaEmbedder struct {
	Model  string
	Client *http.Client
	URL    *url.URL
}

// NewOllamaEmbedder создает новый экземпляр OllamaEmbedder.
func NewOllamaEmbedder(model, ollamaAddress string, timeout time.Duration) (*OllamaEmbedder, error) {
	parsedURL, err := url.Parse(fmt.Sprintf("%s%s", ollamaAddress, embeddingsURL))
	if err != nil {
		return nil, fmt.Errorf("embedding: failed to parse URL, %w", err)
	}
	return &OllamaEmbedder{
		Model:  model,
		Client: &http.Client{Timeout: timeout},
		URL:    parsedURL,
	}, nil
}

// GenerateEmbeddings генерирует эмбеддинговое представление передаваемых текстовых данных.
func (o *OllamaEmbedder) GenerateEmbeddings(ctx context.Context, input string) ([]float32, error) {
	rawBody := map[string]interface{}{
		"model":  o.Model,
		"prompt": input,
	}

	reqBody, err := json.Marshal(rawBody)
	if err != nil {
		return nil, fmt.Errorf("embedding: failed to marshal request, %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, o.URL.Scheme, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("embedding: failed to create request, %w", err)
	}

	resp, err := o.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("embedding: HTTP request failed, %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("embedding: failed to generate embeddings, error %d", resp.StatusCode)
	}

	var response OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("embedding: failed to decode response: %w", err)
	}

	return response.Embedding, nil
}
