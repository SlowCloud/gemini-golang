package main

// promptui 버리고 huh로 바꿔보기

import (
	"fmt"
	"os"
	"time"

	"github.com/SlowCloud/gemini-golang/core"
	"github.com/charmbracelet/huh"
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

	goChat := core.NewGoChatUsecase()

	for {
		var s string

		form := huh.NewText().
			Title("입력").Description("escape: /exit").Value(&s)

		form.Run()

		if s == "" {
			break
		}

		fmt.Println("> " + s)

		if s == "/exit" {
			break
		}

		ch, errCh := goChat.ChatStream(s)

		for tok := range ch {
			fmt.Print(tok)
		}
		if err := <-errCh; err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println()
	}

	history, err := goChat.GetHistory()
	if err != nil {
		fmt.Println("Error getting history:", err)
		return
	}

	now := time.Now().Local().Format("2006-01-02_15:04:05")
	filename := fmt.Sprintf("chat_history-%s.txt", now)

	os.WriteFile(filename, history, 0644)

}
