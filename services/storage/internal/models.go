package internal

// Metadata - структура метаданных документа.
type Metadata struct {
	URL string `json:"url"`
}

// Document структура, определяющая объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	ID        string    // uuid
	Content   string    `json:"content"`  // содержание
	Metadata  Metadata  `json:"metadata"` // метаданные
	Embedding []float32 // векторное представление
	Score     float32   // используется при запросах (queries), показывает схожесть документа с пришедшим в запросе
}
