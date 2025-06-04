package internal

// User определяет структуру пользователя.
type User struct {
	Id         int
	Login      string
	Password   string
	Active     bool
	Collection string
}
