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
	"os/signal"
	"strings"
	"syscall"
	"time"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"google.golang.org/genai"
)

func main() {

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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println()
		fmt.Println("채팅을 종료합니다.")
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		s := strings.TrimSpace(scanner.Text())

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
