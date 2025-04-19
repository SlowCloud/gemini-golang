package ui

import (
	"fmt"
	"log"
	"strings"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const gap = "\n\n"

type (
	errMsg error
)

type Model struct {
	gemini      *gemini.Gemini
	messages    []string
	viewport    viewport.Model
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func New(gemini *gemini.Gemini) Model {
	ta := textarea.New()
	ta.Placeholder = "내용을 입력해주세요."
	ta.Focus()

	ta.Prompt = "| "

	ta.SetWidth(30)
	ta.SetHeight(3)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`gemini 채팅입니다!
자유롭게 메시지를 입력하시고 ctrl+e 를 눌러 대화하세요.`)

	ta.KeyMap.InsertNewline.SetEnabled(true)

	return Model{
		gemini:   gemini,
		messages: []string{},
		viewport: vp,
		textarea: ta,
	}
}

func (m Model) Init() tea.Cmd {
	log.Println("initializing chat model...")
	return tea.Batch(
		textarea.Blink,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		m.viewport.Height = msg.Height - m.textarea.Height() - lipgloss.Height(gap)

		if len(m.messages) > 0 {
			// Wrap content before setting it.
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
		}
		m.viewport.GotoBottom()
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		// case tea.KeyEnter:
		case tea.KeyCtrlE:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m Model) View() string {
	return ""
}

func geminiCmd() tea.Msg {
	return nil
}
