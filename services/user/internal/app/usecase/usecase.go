package usecase

import "github.com/child6yo/rago/services/user/internal/app/repository"

// Usecase содержит все интерфейсы пакета юзкейса.
type Usecase struct {
	Authorization
	APIKey
	Collection
}

// NewUsecase создает новый юзкейс.
func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		Authorization: NewAuthorizationService(repository.Authorization),
		APIKey:        NewAPIKeyService(repository.APIKey),
		Collection:    NewCollectionService(repository.Collection),
	}
}
