package core

import (
	"context"
	"log"
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
		log.Fatal("클라이언트 생성 실패")
		panic(err)
	}
	defer cancel()

	ctx, cancel = context.WithTimeout(background, 10*time.Second)
	chat, err := client.Chats.Create(ctx, "gemini-2.5-flash", nil, nil)
	if err != nil {
		log.Fatal("채팅 세션 생성 실패")
		panic(err)
	}
	defer cancel()

	return goChatUsecase{chat}
}

func (g goChatUsecase) Chat(text string) string {
	panic("unimplemented")
}

// chatStream implements ChatUsecase.
func (g goChatUsecase) ChatStream(text string) (<-chan string, error) {
	// background := context.Background()
	// ctx, cancel := context.WithTimeout(background, 1*time.Minute)
	iter := g.chat.SendMessageStream(context.TODO(), genai.Part{Text: text})
	// defer cancel()

	outputChan := make(chan string)
	go func() {
		defer close(outputChan)
		log.Println("채팅 받는 중...")
		for tok := range iter {
			outputChan <- tok.Text()
		}
		log.Println("채팅 완료.")
	}()

	return outputChan, nil
}

var _ ChatUsecase = goChatUsecase{}
