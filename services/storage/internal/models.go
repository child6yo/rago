package internal

// Metadata - структура метаданных документа.
type Metadata struct {
	URL string `json:"url"`
}

// Document структура, определяющая объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	Content  string   `json:"content"` // содержание
	Metadata Metadata `json:"metadata"` // метаданные
	Score float32 // используется при запросах (queries), показывает схожесть документа с пришедшим в запросе
}
