package core

import (
	"context"
	"os"
	"time"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"google.golang.org/genai"
)

type goChatUsecase struct {
	chat *genai.Chat
}

func NewGoChatUsecase() ChatUsecase {
	background := context.Background()

	ctx, cancel := context.WithTimeout(background, 10*time.Second)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: os.Getenv(configuration.DefaultApiKeyEnviromentVariable)})
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

	return goChatUsecase{chat}
}

func (g goChatUsecase) Chat(text string) string {
	panic("unimplemented")
}

func (g goChatUsecase) ChatStream(text string) (<-chan string, <-chan error) {
	iter := g.chat.SendMessageStream(context.TODO(), genai.Part{Text: text})

	outputChan := make(chan string)
	errChan := make(chan error, 1)
	go func() {
		defer close(outputChan)
		for tok, err := range iter {
			if err != nil {
				errChan <- err
				close(errChan)
				return
			}
			outputChan <- tok.Text()
		}
	}()

	return outputChan, errChan
}

var _ ChatUsecase = goChatUsecase{}
