package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/child6yo/rago/services/user/internal/app/repository"
)

const (
	keyLength = 128
)

// APIKeyService имплементирует интерфейс APIKey.
type APIKeyService struct {
	repo repository.APIKey
}

// NewAPIKeyService создает новый экземпляр APIKeyService.
func NewAPIKeyService(repo repository.APIKey) *APIKeyService {
	return &APIKeyService{repo: repo}
}

// CreateAPIKey создает новый API ключ для пользователя.
func (acs *APIKeyService) CreateAPIKey(userID int) (string, error) {
	key, err := generateAPIKey()
	if err != nil {
		return "", err
	}

	err = acs.repo.CreateAPIKey(userID, key)
	if err != nil {
		return "", err
	}

	return key, nil
}

// DeleteAPIKey удаляет API ключ.
func (acs *APIKeyService) DeleteAPIKey(keyID int, userID int) error {
	return acs.repo.DeleteAPIKey(keyID, userID)
}

// GetAPIKeys возвращает все API ключи пользователя.
func (acs *APIKeyService) GetAPIKeys(userID int) ([]internal.APIKey, error) {
	return acs.repo.GetAPIKeys(userID)
}

// generateAPIKey генерирует API-ключ формата Base64.
func generateAPIKey() (string, error) {
	key := make([]byte, keyLength)

	if _, err := rand.Read(key); err != nil {
		return "", fmt.Errorf("usecase: failed to generate API key: %w", err)
	}

	apiKey := base64.URLEncoding.EncodeToString(key)
	apiKey = strings.TrimRight(apiKey, "=")
	return apiKey, nil
}

// CheckAPIKey валидирует API ключ.
func (acs *APIKeyService) CheckAPIKey(key string) error {
	return acs.repo.CheckAPIKey(key)
}
