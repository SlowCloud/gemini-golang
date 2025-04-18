package gemini_test

import (
	"testing"

	"github.com/SlowCloud/gemini-golang/gemini"
)

func TestGeminiCreate(t *testing.T) {
	gemini, err := gemini.New()
	if err != nil {
		t.Fatal(err)
	}
	if gemini == nil {
		t.FailNow()
	}
}
