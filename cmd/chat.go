/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/SlowCloud/gemini-golang/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "gemini와 채팅을 시작합니다.",
	Long: `gemini와 채팅을 시작합니다. ctrl+e를 눌러 채팅을 전송할 수 있습니다.
채팅 내용은 기록되지 않습니다.
(추후 업데이트)`,
	Run: func(cmd *cobra.Command, args []string) {
		gemini, err := gemini.New()
		if err != nil {
			log.Fatal(err)
		}
		model := ui.New(gemini)

		p := tea.NewProgram(model, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
