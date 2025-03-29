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
	ctx    context.Context
}

func NewGemini() (*Gemini, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return &Gemini{client: client, ctx: ctx}, nil
}

func (gemini *Gemini) Generate(words string) (string, error) {
	result, err := gemini.client.Models.GenerateContent(gemini.ctx, defaultModel, genai.Text(words), nil)
	if err != nil {
		return "", err
	}
	return result.Text(), nil
}

func init() {
	defaultModel = "gemini-2.0-flash-exp"
	apiKey = os.Getenv("GEMINI_API_KEY")
}
