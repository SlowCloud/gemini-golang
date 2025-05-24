package gemini_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/SlowCloud/gemini-golang/gemini"
	"google.golang.org/genai"
)

func TestGeminiCreate(t *testing.T) {
	gemini := createClient()
	if gemini == nil {
		t.Fatal("Failed to create Gemini instance")
	}
}

func TestGeminiAsk(t *testing.T) {
	skipShort(t)
	client := createClient()
	msg, err := gemini.Ask(client, "hello!")
	if err != nil {
		t.Fatalf("Ask failed: %v", err)
	}
	if len(msg) == 0 {
		t.Fail()
	}
	fmt.Println("ask: hello!")
	fmt.Println("answer: " + msg)
}

func TestCreateChatSession(t *testing.T) {
	client := createClient()
	createChatSession(client)
}

func TestChatSessionChat(t *testing.T) {
	skipShort(t)
	client := createClient()
	chatSession := createChatSession(client)
	msg, err := gemini.Chat(chatSession, "hello!")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ask: hello!")
	fmt.Println("res: " + msg)
}

func createClient() *genai.Client {
	client, err := gemini.New()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func skipShort(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}
}

func createChatSession(client *genai.Client) *genai.Chat {
	chat, err := gemini.CreateChatSession(client)
	if err != nil {
		log.Fatal(err)
	}
	return chat
}
