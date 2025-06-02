package usecase

// DocHandler определяет интерфейс обработки пришедших документов.
type DocHandler interface {
	// HandleDocMessage обрабатывает закодированные json-документы,
	// декодирует их в структуры и передает далее в векторную базу данных.
	HandleDocMessage(message []byte) error
}


