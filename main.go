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

		// 텍스트를 밖으로 뱉는 방식이 아니라 참조 방식인 게 좀 그렇네
		// 도움말 띄울 수 있으면 좋을 것 같은데.
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

	chatLoop:
		for {
			select {
			case tok, ok := <-ch:
				if !ok {
					break chatLoop
				}
				fmt.Print(tok)
			case err := <-errCh:
				fmt.Println("Error:", err)
				break chatLoop
			}
		}

		fmt.Println()
	}
}
