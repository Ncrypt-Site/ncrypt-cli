package main

import (
	"github.com/c-bata/go-prompt"
	"testing"
)

func TestBuildSuggestions(t *testing.T) {
	s := buildSuggestions()
	if len(s.noSuggestion(prompt.Document{})) != 0 {
		t.Fail()
	}

	if len(s.destructAfterOpeningSuggestions(prompt.Document{})) != 2 {
		t.Fail()
	}

	if len(s.selfDestructSuggestions(prompt.Document{})) != 9 {
		t.Fail()
	}
}
