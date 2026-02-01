package core

type ChatUsecase interface {
	chat(text string) string
	chatStream(text string) (<-chan string, error)
}
