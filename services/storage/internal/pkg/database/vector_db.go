package database

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

// Все поддерживаемые векторные движки и хранилища должны имплементировать
// нижеописанный интерфейс взаимодействия с ними.

// VectorDB определяет интерфейс взаимодействия с векторным хранилищем данных.
type VectorDB interface {
	// Put позволяет загрузить единицу данных в хранилище.
	Put(ctx context.Context, docs []schema.Document) error

	// Query позволяет выполнить векторный поиск по хранилищу.
	Query(ctx context.Context, query string, numDocs int) ([]schema.Document, error)
}
