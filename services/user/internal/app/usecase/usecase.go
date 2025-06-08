package usecase

import "github.com/child6yo/rago/services/auth/internal/app/repository"

type Usecase struct {
	Authorization
	ApiKey
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		Authorization: NewAuthorizationService(repository.Authorization),
		ApiKey: NewApiKeyService(repository.ApiKey),
	}
}
