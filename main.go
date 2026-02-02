package main

// promptui 버리고 huh로 바꿔보기

import (
	"fmt"

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
}
