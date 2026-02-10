package core

type Repository interface {
	GetHistoryList() ([]string, error)
	LoadHistory(filename string) ([]Message, error)
	SaveHistory(filename string, history []Message) error
}

type Message struct {
	Message string
	Role    string
}
