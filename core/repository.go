package core

type Repository[T any] interface {
	GetHistoryList() ([]string, error)
	LoadHistory(filename string) (T, error)
	SaveHistory(filename string, history T) error
}
