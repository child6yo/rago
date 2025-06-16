package internal

// Metadata - структура метаданных документа.
type Metadata struct {
	URL string `json:"url"`
}

// Document структура, определяющая объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	Content    string   `json:"content"`
	Metadata   Metadata `json:"metadata"`
	Collection string   `json:"collection"`
}

// DocumentArray определяет массив документов с указанием коллекции,
// к которой они относятся.
type DocumentArray struct {
	Documents  []Document `json:"documents"`
	Collection string     `json:"collection"` // к какой коллекции относятся документы
}
