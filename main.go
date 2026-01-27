/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"github.com/manifoldco/promptui"
	"google.golang.org/genai"
)

func main() {

menuLoop:
	for {

		prompt := promptui.Select{
			Label: "This is Menu",
			Items: []string{
				"start chat",
				"select chat",
				"exit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		switch result {
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

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		s := strings.TrimSpace(scanner.Text())

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
