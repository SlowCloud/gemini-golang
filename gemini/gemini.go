package gemini

import (
	"context"
	"os"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
}

func New() (*Gemini, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv(configuration.DefaultApiKeyEnviromentVariable),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return &Gemini{client: client}, nil
}

func (gemini *Gemini) Ask(words string) (string, error) {
	ctx := context.Background()
	result, err := gemini.client.Models.GenerateContent(ctx, configuration.DefaultGeminiModel, genai.Text(words), nil)
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
	chat, err := gemini.client.Chats.Create(ctx, configuration.DefaultGeminiModel, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatSession{chat: chat}, nil
}
