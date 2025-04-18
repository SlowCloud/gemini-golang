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

func TestGeminiCreate(t *testing.T) {
	gemini := createGemini()
	if gemini == nil {
		t.FailNow()
	}
}

func TestGeminiAsk(t *testing.T) {
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
