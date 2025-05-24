package gemini

import (
	"context"
	"os"

	configuration "github.com/SlowCloud/gemini-golang/config"
	"google.golang.org/genai"
)

func New() (*genai.Client, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv(configuration.DefaultApiKeyEnviromentVariable),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Ask(client *genai.Client, words string) (string, error) {
	ctx := context.Background()
	result, err := client.Models.GenerateContent(ctx, configuration.DefaultGeminiModel, genai.Text(words), nil)
	if err != nil {
		return "", err
	}
	return result.Text(), nil
}

func Chat(chat *genai.Chat, text string) (string, error) {
	ctx := context.Background()
	res, err := chat.SendMessage(ctx, genai.Part{Text: text})
	if err != nil {
		return "", err
	}
	return res.Text(), nil
}

func CreateChatSession(client *genai.Client) (*genai.Chat, error) {
	ctx := context.Background()
	chat, err := client.Chats.Create(ctx, configuration.DefaultGeminiModel, nil, nil)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
