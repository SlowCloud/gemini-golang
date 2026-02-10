package main

// promptui 버리고 huh로 바꿔보기

import (
	"fmt"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/charmbracelet/huh"
)

func main() {

	goChat := gemini.NewGoChatUsecase()

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
