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
