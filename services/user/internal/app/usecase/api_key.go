package usecase

import (
	"github.com/child6yo/rago/services/user/internal"
	"github.com/child6yo/rago/services/user/internal/app/repository"
)

type ApiKeyService struct {
	repo repository.ApiKey
}

func NewApiKeyService(repo repository.ApiKey) *ApiKeyService {
	return &ApiKeyService{repo: repo}
}

func (acs *ApiKeyService) CreateApiKey(userID int, key string) error {
	return acs.repo.CreateApiKey(userID, key)
}

func (acs *ApiKeyService) DeleteApiKey(keyID int, userID int) error {
	return acs.repo.DeleteApiKey(keyID, userID)
}

func (acs *ApiKeyService) GetApiKeys(userID int) ([]internal.ApiKey, error) {
	return acs.repo.GetApiKeys(userID)
}
