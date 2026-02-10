package gemini

import (
	"context"
	"time"

	"github.com/SlowCloud/gemini-golang/core"
	"google.golang.org/genai"
)

type geminiChatUsecase struct {
	chat *genai.Chat
}

func NewGoChatUsecase() core.ChatUsecase {
	background := context.Background()

	ctx, cancel := context.WithTimeout(background, 10*time.Second)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{})
	if err != nil {
		panic(err)
	}
	defer cancel()

	ctx, cancel = context.WithTimeout(background, 10*time.Second)
	chat, err := client.Chats.Create(ctx, "gemini-2.5-flash", nil, nil)
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

var _ core.ChatUsecase = geminiChatUsecase{}
