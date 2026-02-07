package main

// promptui 버리고 huh로 바꿔보기

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SlowCloud/gemini-golang/core"
	"github.com/charmbracelet/huh"
)

func main() {

	actions := map[string]func(){
		"start chat": chat,
	}

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

		if selected == "exit" {
			break
		}

		if action, ok := actions[selected]; ok {
			action()
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

	err := saveHistory(goChat)
	if err != nil {
		panic(err)
	}

}

func saveHistory(goChat core.ChatUsecase) error {
	history, err := goChat.GetHistory()
	if err != nil {
		fmt.Println("Error getting history:", err)
		return err
	}

	now := time.Now().Local().Format("2006-01-02_150405")
	filename := fmt.Sprintf("chat_history-%s.txt", now)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}
	err = os.WriteFile(filepath.Join(dir, filename), history, 0644)
	if err != nil {
		fmt.Println("Error writing history to file:", err)
		return err
	}

	fmt.Println("Chat history saved to", filename, "path ", dir)
	return nil
}
