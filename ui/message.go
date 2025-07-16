package ui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/genai"
)

type Message struct {
	content   string
	streaming bool
	channel   chan string
	chat      *genai.Chat
}

type streamMsg struct {
	token   string
	isFinal bool
}

type streamEndMsg struct{}

func NewMessage(chat *genai.Chat, message string) *Message {
	return &Message{
		content:   "",
		streaming: false,
		chat:      chat,
		channel:   make(chan string),
	}
}

func (m *Message) Init() tea.Cmd {
	go m.startStream()
	return m.streamNextTokenCmd()
}

func (m *Message) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case streamMsg:
		m.content += msg.token
		if msg.isFinal {
			m.streaming = false
			return m, func() tea.Msg { return streamEndMsg{} }
		}
		return m, m.streamNextTokenCmd()
	case streamEndMsg:
		m.streaming = false
		return m, nil
	}
	return m, nil
}

func (m *Message) View() string {
	return m.content
}

func (m *Message) streamNextTokenCmd() tea.Cmd {
	return func() tea.Msg {
		token, ok := <-m.channel
		if ok {
			return streamMsg{token: token, isFinal: false}
		}
		return streamMsg{token: "", isFinal: true}
	}
}

func (m *Message) startStream() {
	m.streaming = true
	for token := range m.chat.SendMessageStream(context.Background(), genai.Part{Text: m.content}) {
		m.channel <- token
	}
	m.streaming = false
	close(m.channel)
}