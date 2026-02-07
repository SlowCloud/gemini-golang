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

			chat(historyData)
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
	if history == nil {
		fmt.Println("No history to save")
		return nil
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

func getHistoryList() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var historyFiles []string

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" && len(file.Name()) >= 12 && file.Name()[:12] == "chat_history" {
			historyFiles = append(historyFiles, file.Name())
		}
	}
	return historyFiles, nil
}

func loadHistory(filename string) ([]byte, error) {
	wd, err := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, filename))
	if err != nil {
		return nil, err
	}
	return data, nil
}
