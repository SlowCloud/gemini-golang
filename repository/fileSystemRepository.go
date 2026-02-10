package repository

import (
	"encoding/json"
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

type FileSystemRepository[T any] struct {
}

func (f FileSystemRepository[T]) GetHistoryList() ([]string, error) {
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

func (f FileSystemRepository[T]) LoadHistory(filename string) (*T, error) {
	wd, err := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, filename))
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (f FileSystemRepository[T]) SaveHistory(filename string, history *T) error {
	if history == nil {
		fmt.Println("No history to save")
		return nil
	}

	bytes, ok := any(history).([]byte)
	if !ok {
		return fmt.Errorf("invalid history type, expected []byte")
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
	err = os.WriteFile(filepath.Join(dir, filename), bytes, 0644)
	if err != nil {
		fmt.Println("Error writing history to file:", err)
		return err
	}

	fmt.Println("Chat history saved to", filename, "path ", dir)
	return nil
}

var _ core.Repository[any] = FileSystemRepository[any]{}
