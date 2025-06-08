package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strings"

	"github.com/child6yo/rago/services/user/internal"
	"github.com/child6yo/rago/services/user/internal/app/repository"
)

const (
	keyLenght = 128
)

type ApiKeyService struct {
	repo repository.ApiKey
}

func NewApiKeyService(repo repository.ApiKey) *ApiKeyService {
	return &ApiKeyService{repo: repo}
}

func (acs *ApiKeyService) CreateApiKey(userID int) (string, error) {
	key, err := generateAPIKey()
	if err != nil {
		log.Print("usecase:", err)
		return "", err
	}

	err = acs.repo.CreateApiKey(userID, key)
	if err != nil {
		log.Print("usecase:", err)
		return "", err
	}

	return key, nil
}

func (acs *ApiKeyService) DeleteApiKey(keyID int, userID int) error {
	return acs.repo.DeleteApiKey(keyID, userID)
}

func (acs *ApiKeyService) GetApiKeys(userID int) ([]internal.ApiKey, error) {
	return acs.repo.GetApiKeys(userID)
}

// generateAPIKey генерирует API-ключ формата Base64.
// TODO - сменить формат (этот не парсится с юрла)
func generateAPIKey() (string, error) {
	key := make([]byte, keyLenght)

	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	apiKey := base64.StdEncoding.EncodeToString(key)
	apiKey = strings.TrimRight(apiKey, "=")
	return apiKey, nil
}

func (acs *ApiKeyService) CheckAPIKey(key string) error {
	return acs.repo.CheckAPIKey(key)
}
