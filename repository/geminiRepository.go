package repository

import (
	"github.com/SlowCloud/gemini-golang/core"
	"google.golang.org/genai"
)

type GeminiRepository interface {
	core.Repository[[]*genai.Content]
}
