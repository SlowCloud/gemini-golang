package gemini_test

import (
	"fmt"
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

func skipShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skip " + t.Name() + " in short mode")
	}
}

func TestGeminiCreate(t *testing.T) {
	gemini := createGemini()
	if gemini == nil {
		t.FailNow()
	}
}

func TestGeminiAsk(t *testing.T) {
	skipShort(t)
	gemini := createGemini()
	msg, err := gemini.Ask("hello!")
	if err != nil {
		log.Fatal(err)
	}
	if len(msg) == 0 {
		t.Fail()
	}
	fmt.Println("ask: hello!")
	fmt.Println("answer: " + msg)
}
