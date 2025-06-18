package internal

// User определяет модель пользователя.
type User struct {
	ID         int
	Login      string
	Password   string
	Active     bool
	Collection string
}

// APIKey определяет модель ключа API.
type APIKey struct {
	ID  int    `db:"id"`
	Key string `db:"key"`
}
