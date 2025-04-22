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

const gap = "\n"

type (
	errMsg error
)

var (
	userChatRenderer lipgloss.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
)

type Model struct {
	gemini      *gemini.Gemini
	session     *gemini.ChatSession
	messages    []string
	viewport    viewport.Model
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func New(gemini *gemini.Gemini) Model {
	// textarea
	ta := newTextarea()
	// viewport
	vp := newViewport()

	// gemini session sttings
	session, err := gemini.CreateChat()
	if err != nil {
		log.Fatal(err)
	}

	return Model{
		gemini:   gemini,
		session:  session,
		messages: []string{},
		viewport: vp,
		textarea: ta,
	}
}

func newViewport() viewport.Model {
	vp := viewport.New(30, 5)

	vp.Style = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder())
	return vp
}

func newTextarea() textarea.Model {
	ta := textarea.New()

	// textarea style
	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder())
	ta.ShowLineNumbers = false

	// textarea sttings
	ta.Focus()
	ta.SetWidth(30)
	ta.SetHeight(3)
	ta.Placeholder = "내용을 입력해주세요."
	ta.Prompt = "  "
	ta.Placeholder = `gemini 채팅입니다!
자유롭게 메시지를 입력하시고 ctrl+e 를 눌러 대화하세요.`
	ta.KeyMap.InsertNewline.SetEnabled(true)

	return ta
}

func (m Model) Init() tea.Cmd {
	log.Println("initializing chat model...")
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	cmds := []tea.Cmd{tiCmd, vpCmd}

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
			text := m.textarea.Value()
			m.messages = append(m.messages, userChatRenderer.Render(m.senderStyle.Render("You: ")+text))
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			m.textarea.Reset()
			m.viewport.GotoBottom()
			m.textarea.Blur()
			cmds = append(cmds, m.createGeminiCmd(text))
		}
	case geminiCmd:
		text := string(msg)
		m.messages = append(m.messages, m.senderStyle.Render("Gemini: ")+text)
		m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
		if !m.textarea.Focused() {
			m.textarea.Focus()
		}
		m.viewport.GotoBottom()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s%s%s",
		m.viewport.View(),
		gap,
		m.textarea.View(),
	)
}

type geminiCmd string

func (m Model) createGeminiCmd(text string) tea.Cmd {
	return func() tea.Msg {
		msg, err := m.session.Chat(text)
		if err != nil {
			log.Fatal(err)
		}
		return geminiCmd(msg)
	}
}
