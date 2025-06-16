package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	embeddingsURL = "/api/embeddings"
	maxRetries    = 3
	retryDelay    = 1 * time.Second
)

// OllamaResponse определяет ответ ollama.
type OllamaResponse struct {
	Embedding []float32 `json:"embedding"`
}

// OllamaEmbedder имплементирует интерфейс Embedder.
type OllamaEmbedder struct {
	Model  string
	Client *http.Client
	URL    string
}

// NewOllamaEmbedder создает новый экземпляр OllamaEmbedder.
func NewOllamaEmbedder(model, ollamaAddress string, timeout time.Duration) (*OllamaEmbedder, error) {
	url := fmt.Sprintf("%s%s", ollamaAddress, embeddingsURL)

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     90 * time.Second,
			DisableCompression:  false,
		},
	}

	return &OllamaEmbedder{
		Model:  model,
		Client: client,
		URL:    url,
	}, nil
}

// GenerateEmbeddings генерирует эмбеддинговое представление передаваемых текстовых данных.
func (o *OllamaEmbedder) GenerateEmbeddings(ctx context.Context, input string) ([]float32, error) {
	rawBody := map[string]interface{}{
		"model":  o.Model,
		"prompt": input,
	}

	var response OllamaResponse
	var lastErr error

	// retry loop
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			truncatedInput := input
			if len(input) > 50 {
				truncatedInput = input[:50] + "..."
			}
			log.Printf("[WARN] Retrying request (attempt %d) for model %s on input: %s",
				attempt, o.Model, truncatedInput)

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(retryDelay * time.Duration(attempt)):
			}
		}

		bodyBytes, err := json.Marshal(rawBody)
		if err != nil {
			lastErr = fmt.Errorf("embedding: marshal request failed: %w", err)
			continue
		}

		req, err := http.NewRequestWithContext(ctx, "POST", o.URL, bytes.NewReader(bodyBytes))
		if err != nil {
			lastErr = fmt.Errorf("embedding: create request failed: %w", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		start := time.Now()
		resp, err := o.Client.Do(req)
		duration := time.Since(start)

		if err != nil {
			log.Printf("[ERROR] HTTP request failed: %v (duration: %v)", err, duration)

			if isNetworkError(err) {
				lastErr = err
				continue
			}
			return nil, fmt.Errorf("embedding: HTTP request failed: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
			resp.Body.Close()
			lastErr = fmt.Errorf("embedding: unexpected status %d: %s", resp.StatusCode, string(body))

			if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				log.Printf("[WARN] Server error (status %d), retrying...", resp.StatusCode)
				continue
			}
			return nil, lastErr
		}

		limitedReader := io.LimitReader(resp.Body, 10*1024*1024)
		if err := json.NewDecoder(limitedReader).Decode(&response); err != nil {
			resp.Body.Close()
			lastErr = fmt.Errorf("embedding: decode failed: %w", err)
			continue
		}
		resp.Body.Close()

		log.Printf("[INFO] Embedding generated in %v, size: %d", duration, len(response.Embedding))

		return response.Embedding, nil
	}

	return nil, fmt.Errorf("embedding: after %d attempts: %w", maxRetries, lastErr)
}

func isNetworkError(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}
