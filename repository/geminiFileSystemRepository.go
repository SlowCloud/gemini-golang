package repository

import (
	"github.com/SlowCloud/gemini-golang/core"
	"google.golang.org/genai"
)

type GeminiFileSystemRepository struct {
	FileSystemRepository[[]*genai.Content]
}

var _ core.Repository[[]*genai.Content] = GeminiFileSystemRepository{}
