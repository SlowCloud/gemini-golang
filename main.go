package main

// promptui 버리고 huh로 바꿔보기

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"github.com/charmbracelet/huh"
	"google.golang.org/genai"
)

func main() {

menuLoop:
	for {

		var selected string
		form := huh.NewSelect[string]().
			Title("Gemini-Golang Menu").
			Options(
				huh.NewOption("start chat", "start chat"),
				huh.NewOption("select chat", "select chat"),
				huh.NewOption("exit", "exit"),
			).
			Value(&selected)

		form.Run()

		switch selected {
		case "start chat":
			chat()
		case "exit":
			break menuLoop
		}

	}

}

func chat() {
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

	for {

		var s string

		// 텍스트를 밖으로 뱉는 방식이 아니라 참조 방식인 게 좀 그렇네
		// 도움말 띄울 수 있으면 좋을 것 같은데.
		form := huh.NewText().
			Title("입력").Description("escape: /exit").Value(&s)

		form.Run()

		fmt.Println("> " + s)

		if s == "/exit" {
			break
		}

		ctx, cancel = context.WithTimeout(background, 1*time.Minute)
		iter := chat.SendMessageStream(ctx, genai.Part{Text: s})

		for tok, err := range iter {
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(tok.Text())
		}

		cancel()

		fmt.Println()

	}
}
