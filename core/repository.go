package core

type Repository interface {
	getHistoryList() ([]string, error)
	loadHistory(filename string) ([]byte, error)
	saveHistory(filename string, history []byte) error
}
