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
	"github.com/manifoldco/promptui"
	"google.golang.org/genai"
)

func main() {

menuLoop:
	for {

		// prompt := promptui.Select{
		// 	Label: "This is Menu",
		// 	Items: []string{
		// 		"start chat",
		// 		"select chat",
		// 		"exit",
		// 	},
		// }
		// _, result, err := prompt.Run()
		// if err != nil {
		// 	panic(err)
		// }
		// switch result {
		// case "start chat":
		// 	chat()
		// case "exit":
		// 	break menuLoop
		// }

		var selected string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Gemini-Golang Menu").
					Options(
						huh.NewOption("start chat", "start chat"),
						huh.NewOption("select chat", "select chat"),
						huh.NewOption("exit", "exit"),
					).
					Value(&selected),
			),
		)

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

		prompt := promptui.Prompt{}
		s, err := prompt.Run()
		if err != nil {
			panic(err)
		}

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
