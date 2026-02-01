package core

type ChatUsecase interface {
	Chat(text string) string
	ChatStream(text string) (<-chan string, error)
}
