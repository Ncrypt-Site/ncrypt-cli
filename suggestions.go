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
			{Text: "1h", Description: "In 1 hour your secure note will be deleted"},
			{Text: "3h", Description: "In 3 hours your secure note will be deleted"},
			{Text: "6h", Description: "In 6 hours your secure note will be deleted"},
			{Text: "12h", Description: "In 12 hours your secure note will be deleted"},
			{Text: "1d", Description: "In 1 day your secure note will be deleted"},
			{Text: "2d", Description: "In 2 days your secure note will be deleted"},
			{Text: "3d", Description: "In 3 days your secure note will be deleted"},
			{Text: "7d", Description: "In 7 days your secure note will be deleted"},
			{Text: "1m", Description: "In 1 month your secure note will be deleted"},
		},
		DestructAfterOpeningSuggestions: []prompt.Suggest{
			{Text: "yes", Description: "your note will be destroyed from our server on the first opening"},
			{Text: "no", Description: "we will preserve your note until it expires"},
		},
	}
}

func (c *Completer) noSuggestion(d prompt.Document) []prompt.Suggest {
	return c.NoSuggestion
}

func (c *Completer) selfDestructSuggestions(d prompt.Document) []prompt.Suggest {
	return prompt.FilterContains(c.NoteSelfDestructSuggestion, d.GetWordBeforeCursor(), true)
}

func (c *Completer) destructAfterOpeningSuggestions(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.DestructAfterOpeningSuggestions, d.GetWordBeforeCursor(), true)
}
