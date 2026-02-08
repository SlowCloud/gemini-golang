package core

type Repository interface {
	GetHistoryList() ([]string, error)
	LoadHistory(filename string) ([]byte, error)
	SaveHistory(filename string, history []byte) error
}
