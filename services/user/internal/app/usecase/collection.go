package usecase

import "github.com/child6yo/rago/services/user/internal/app/repository"

// CollectionService имплементирует интерфейс Collection.
type CollectionService struct {
	repo repository.Collection
}

// NewCollectionService создает новый экземпляр CollectionService.
func NewCollectionService(repo repository.Collection) *CollectionService {
	return &CollectionService{repo: repo}
}

// GetCollection возвращает коллекцию пользователя.
func (cs *CollectionService) GetCollection(userID int) (string, error) {
	return cs.repo.GetCollection(userID)
}
