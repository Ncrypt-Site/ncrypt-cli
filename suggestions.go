package main

import "github.com/c-bata/go-prompt"

type Completer struct {
	NoSuggestion                    []prompt.Suggest
	NoteSelfDestructSuggestion      []prompt.Suggest
	DestructAfterOpeningSuggestions []prompt.Suggest
}

func buildSuggestions() Completer {
	return Completer{
		NoSuggestion: []prompt.Suggest{},
		NoteSelfDestructSuggestion: []prompt.Suggest{
			{Text: "1 hour", Description: "The time your note will live on our servers"},
			{Text: "3 hours", Description: "The time your note will live on our servers"},
			{Text: "6 hours", Description: "The time your note will live on our servers"},
			{Text: "12 hours", Description: "The time your note will live on our servers"},
			{Text: "1 day", Description: "The time your note will live on our servers"},
			{Text: "2 days", Description: "The time your note will live on our servers"},
			{Text: "3 days", Description: "The time your note will live on our servers"},
			{Text: "7 days", Description: "The time your note will live on our servers"},
			{Text: "1 month", Description: "The time your note will live on our servers"},
		},
		DestructAfterOpeningSuggestions: []prompt.Suggest{
			{Text: "yes", Description: "your note will be destroyed from our server on the first opening"},
			{Text: "no", Description: "we will preserve your note untill it expires"},
		},
	}
}

func (c *Completer) noSuggestion(d prompt.Document) []prompt.Suggest {
	return c.NoSuggestion
}

func (c *Completer) selfDestructSuggestions(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.NoteSelfDestructSuggestion, d.GetWordBeforeCursor(), true)
}

func (c *Completer) destructAfterOpeningSuggestions(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.DestructAfterOpeningSuggestions, d.GetWordBeforeCursor(), true)
}
