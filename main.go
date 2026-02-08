package main

// promptui 버리고 huh로 바꿔보기

import (
	"fmt"

	"github.com/SlowCloud/gemini-golang/core"
	"github.com/SlowCloud/gemini-golang/repository"
	"github.com/charmbracelet/huh"
)

var repo core.Repository[[]byte] = repository.FileSystemRepository[[]byte]{}

func main() {

	actions := map[string]func(){
		"start chat": func() { chat(nil) },
		"select chat": func() {
			historyFiles, err := getHistoryList()
			if err != nil {
				fmt.Println("Error getting history list:", err)
				return
			}
			if len(historyFiles) == 0 {
				fmt.Println("No chat history files found.")
				return
			}

			historyOptions := make([]huh.Option[string], len(historyFiles))
			for i, file := range historyFiles {
				historyOptions[i] = huh.NewOption(file, file)
			}
			var selected string
			form := huh.NewSelect[string]().
				Title("Select Chat History").
				Options(historyOptions...).
				Value(&selected)

			form.Run()

			historyData, err := loadHistory(selected)
			if err != nil {
				fmt.Println("Error loading history:", err)
				return
			}

			chat(*historyData)
		},
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

func chat(history []byte) {

	var goChat core.ChatUsecase
	if history == nil {
		goChat = core.NewGoChatUsecase()
	} else {
		goChat = core.NewGoChatUsecaseWithHistory(history)
	}

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
		panic(err)
	}
	err = saveHistory(history)
	if err != nil {
		panic(err)
	}

}

func saveHistory(history []byte) error {
	return repo.SaveHistory("", &history)
}

func getHistoryList() ([]string, error) {
	return repo.GetHistoryList()
}

func loadHistory(filename string) (*[]byte, error) {
	return repo.LoadHistory(filename)
}
