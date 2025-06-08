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

// Document структура, определяющая объект, который необходимо
// содержать в векторной базе данных.
type Document struct {
	Content  string   `json:"content"`
	Metadata Metadata `json:"metadata"`
}
