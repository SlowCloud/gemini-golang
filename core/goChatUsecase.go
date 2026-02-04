package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/genai"
)

type goChatUsecase struct {
	chat *genai.Chat
}

func (g goChatUsecase) SaveHistory() error {
	history := g.chat.History(false)

	b, err := json.Marshal(history)
	if err != nil {
		return err
	}

	now := time.Now().Local().Format("2006-01-02_150405")
	filename := fmt.Sprintf("chat_history-%s.txt", now)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}

	err = os.WriteFile(filepath.Join(dir, filename), b, 0644)
	if err != nil {
		fmt.Println("Error writing history to file:", err)
		return err
	}

	fmt.Println("Chat history saved to", filename, "path ", dir)

	return nil
}

func NewGoChatUsecase() ChatUsecase {
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

var _ ChatUsecase = goChatUsecase{}
