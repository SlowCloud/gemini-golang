package gemini

import (
	"context"
	"os"

	"google.golang.org/genai"
)

var (
	defaultModel string
	apiKey       string
)

type Gemini struct {
	client *genai.Client
}

func New() (*Gemini, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return &Gemini{client: client}, nil
}

func (gemini *Gemini) Ask(words string) (string, error) {
	ctx := context.Background()
	result, err := gemini.client.Models.GenerateContent(ctx, defaultModel, genai.Text(words), nil)
	if err != nil {
		return "", err
	}
	return result.Text(), nil
}

type ChatSession struct {
	chat *genai.Chat
}

func (c *ChatSession) Chat(text string) (string, error) {
	ctx := context.Background()
	res, err := c.chat.SendMessage(ctx, genai.Part{Text: text})
	if err != nil {
		return "", err
	}
	return res.Text(), nil
}

func (gemini *Gemini) CreateChat() (*ChatSession, error) {
	ctx := context.Background()
	chat, err := gemini.client.Chats.Create(ctx, defaultModel, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatSession{chat: chat}, nil
}

func init() {
	defaultModel = "gemini-2.0-flash-exp"
	apiKey = os.Getenv("GEMINI_API_KEY")
}
