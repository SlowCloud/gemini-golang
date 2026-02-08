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
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(dir)
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
	wd, err := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, filename))
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

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}
	err = os.WriteFile(filepath.Join(dir, filename), history, 0644)
	if err != nil {
		fmt.Println("Error writing history to file:", err)
		return err
	}

	fmt.Println("Chat history saved to", filename, "path ", dir)
	return nil
}

var _ core.Repository = (*FileSystemRepository)(nil)
