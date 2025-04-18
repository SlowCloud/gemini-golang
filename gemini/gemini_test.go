package gemini_test

import (
	"log"
	"testing"

	"github.com/SlowCloud/gemini-golang/gemini"
)

func createGemini() *gemini.Gemini {
	gemini, err := gemini.New()
	if err != nil {
		log.Fatal(err)
	}
	return gemini
}

func TestGeminiCreate(t *testing.T) {
	gemini := createGemini()
	if gemini == nil {
		t.FailNow()
	}
}
