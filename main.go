/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/SlowCloud/gemini-golang/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	gemini, err := gemini.New()
	if err != nil {
		log.Fatal(err)
	}
	model := ui.New(gemini)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
