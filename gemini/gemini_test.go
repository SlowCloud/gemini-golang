package gemini_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/SlowCloud/gemini-golang/gemini"
)

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
		t.Fatal(err)
	}
	if len(msg) == 0 {
		t.Fail()
	}
	fmt.Println("ask: hello!")
	fmt.Println("answer: " + msg)
}

func TestCreateChatSession(t *testing.T) {
	gemini := createGemini()
	createChatSession(gemini)
}

func TestChatSessionChat(t *testing.T) {
	skipShort(t)
	gemini := createGemini()
	chatSession := createChatSession(gemini)
	msg, err := chatSession.Chat("hello!")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ask: hello!")
	fmt.Println("res: " + msg)
}

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

func createChatSession(gemini *gemini.Gemini) *gemini.ChatSession {
	chat, err := gemini.CreateChat()
	if err != nil {
		log.Fatal(err)
	}
	return chat
}
