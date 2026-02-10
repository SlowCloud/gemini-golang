package core

import (
	"context"
	"encoding/json"
	"time"

	"google.golang.org/genai"
)

type geminiChatUsecase struct {
	chat *genai.Chat
}

func (g geminiChatUsecase) GetHistory() ([]byte, error) {
	history := g.chat.History(false)

	b, err := json.Marshal(history)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func NewGoChatUsecase() ChatUsecase {
	return NewGoChatUsecaseWithHistory(nil)
}

func NewGoChatUsecaseWithHistory(history []byte) ChatUsecase {
	background := context.Background()

	ctx, cancel := context.WithTimeout(background, 10*time.Second)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{})
	if err != nil {
		panic(err)
	}
	defer cancel()

	var h []*genai.Content = nil
	if history != nil {
		err = json.Unmarshal(history, &h)
		if err != nil {
			panic(err)
		}
	}

	ctx, cancel = context.WithTimeout(background, 10*time.Second)
	chat, err := client.Chats.Create(ctx, "gemini-2.5-flash", nil, h)
	if err != nil {
		panic(err)
	}
	defer cancel()

	return geminiChatUsecase{chat}
}

func (g geminiChatUsecase) Chat(text string) string {
	panic("unimplemented")
}

func (g geminiChatUsecase) ChatStream(text string) (<-chan string, <-chan error) {
	iter := g.chat.SendMessageStream(context.TODO(), genai.Part{Text: text})

	outputChan := make(chan string)
	errChan := make(chan error, 1)
	go func() {
		defer close(outputChan)
		defer close(errChan)
		for tok, err := range iter {
			if err != nil {
				errChan <- err
				return
			}
			outputChan <- tok.Text()
		}
	}()

	return outputChan, errChan
}

var _ ChatUsecase = geminiChatUsecase{}
