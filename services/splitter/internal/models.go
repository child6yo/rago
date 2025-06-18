package internal

// Metadata - структура метаданных документа.
type Metadata struct {
	URL string `json:"url"`
}

// Document структура, определяющая объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	Content    string   `json:"content,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Collection string   `json:"collection,omitempty"`
}

// DocumentArray определяет массив документов с указанием коллекции,
// к которой они относятся.
type DocumentArray struct {
	Documents  []Document `json:"documents,omitempty"`
	Collection string     `json:"collection,omitempty"` // к какой коллекции относятся документы
}
