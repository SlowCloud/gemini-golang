package repository

import (
	"github.com/SlowCloud/gemini-golang/core"
	"google.golang.org/genai"
)

type GoFileSystemRepository struct {
	FileSystemRepository[[]*genai.Content]
}

var _ core.Repository[[]*genai.Content] = GoFileSystemRepository{}
