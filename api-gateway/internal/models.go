package internal

type User struct {
	Id         int
	Login      string `json:"login"`
	Password   string `json:"password"`
	Active     bool
	Collection string
}

// ApiKey определяет модель ключа API.
type ApiKey struct {
	ID  int
	Key string
}

// Metadata - структура метаданных документа.
type Metadata struct {
	URL string `json:"url"`
}

// Document определяет объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	ID       string   `json:"id"`       // uuid
	Content  string   `json:"content"`  // содержание
	Metadata Metadata `json:"metadata"` // метаданные
}

// DocumentArray определяет массив документов с указанием коллекции,
// к которой они относятся.
type DocumentArray struct {
	Documents  []Document `json:"documents"`
	Collection string     `json:"collection"` // к какой коллекции относятся документы
}
