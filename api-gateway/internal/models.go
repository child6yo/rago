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
