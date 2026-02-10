package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SlowCloud/gemini-golang/core"
)

const (
	historyDir      = "histories"
	historyFilePref = "chat_history_"
	historyFileExt  = ".txt"
)

type FileSystemRepository struct {
}

func (f FileSystemRepository) GetHistoryList() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	targetDir := filepath.Join(homeDir, historyDir)
	files, err := os.ReadDir(targetDir)
	if err != nil {
		return nil, err
	}
	var historyFiles []string

	for _, file := range files {
		if isHistoryFile(file) {
			historyFiles = append(historyFiles, file.Name())
		}
	}
	return historyFiles, nil
}

func isHistoryFile(file os.DirEntry) bool {
	return !file.IsDir() && filepath.Ext(file.Name()) == historyFileExt &&
		len(file.Name()) >= len(historyFilePref) &&
		file.Name()[:len(historyFilePref)] == historyFilePref
}

func (f FileSystemRepository) LoadHistory(filename string) ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	targetDir := filepath.Join(homeDir, historyDir)
	data, err := os.ReadFile(filepath.Join(targetDir, filename))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f FileSystemRepository) SaveHistory(filename string, history []byte) error {
	if history == nil {
		fmt.Println("No history to save")
		return nil
	}

	now := time.Now().Local().Format("2006-01-02_150405")
	if filename == "" {
		filename = fmt.Sprintf("%s-%s%s", historyFilePref, now, historyFileExt)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}
	targetDir := filepath.Join(homeDir, historyDir)
	err = os.WriteFile(filepath.Join(targetDir, filename), history, 0644)
	if err != nil {
		fmt.Println("Error writing history to file:", err)
		return err
	}

	fmt.Println("Chat history saved to", filename, "path ", targetDir)
	return nil
}

var _ core.Repository = FileSystemRepository{}
