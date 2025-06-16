package usecase

// Splitter определяет интерфейс разделителя крупного массива документов
// на единичные документы.
type Splitter interface {
	// SplitDocuments разбивает массив документов на единичные структуры
	// и асинхронно отправляет их в нижележащий сервис.
	SplitDocuments(docs []byte) error
}
