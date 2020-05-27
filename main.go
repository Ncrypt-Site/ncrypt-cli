package main

import (
	"encoding/json"
	"github.com/c-bata/go-prompt"
	"github.com/gookit/color"
	"ncrypt-cli/helpers"
	"ncrypt-cli/models"
	"ncrypt-cli/services"
	"net/http"
	"os"
)

func main() {
	c := buildSuggestions()

	color.FgLightBlue.Println("nCrypt.site CLI V1.0.0-alpha\n")

	color.Green.Print("enter your note:")
	color.Cyan.Println(" (we will secure it for you)")
	note := prompt.Input("> ", c.noSuggestion)

	color.Green.Println("how long should we retain the note?")
	selfDestructInput := prompt.Input("> ", c.selfDestructSuggestions)
	selfDestruct, err := helpers.ConvertSelfDestructToInt(selfDestructInput)
	if err != nil {
		color.Red.Println(err)
		os.Exit(2)
	}

	color.Green.Println("should we destroy the secure note once it's opened?")
	destructAfterOpeningInput := prompt.Input("> ", c.destructAfterOpeningSuggestions)
	destructAfterOpening, err := helpers.ConvertDestructAfterOpeningToBool(destructAfterOpeningInput)
	if err != nil {
		color.Red.Println(err)
		os.Exit(2)
	}

	encryptedNote, err := services.CreateNote([]byte(note))
	if err != nil {
		color.Red.Println("an error occurred, please try again")
		os.Exit(2)
	}

	model := models.SecureMessageRequest{
		Note:                 encryptedNote.Note,
		SelfDestruct:         selfDestruct,
		DestructAfterOpening: destructAfterOpening,
	}

	h := services.HttpService{
		Client: &http.Client{},
		Url:    "https://api.ncrypt.site/api/v1/note",
	}

	resp, err := h.SendRequest(model)
	if err != nil {
		color.Red.Println(err)
		os.Exit(2)
	}

	r := models.NoteCreatedResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		color.Red.Println(err)
		os.Exit(2)
	}

	color.Green.Println("Note has been created, the following url is now holds your secure note:")
	color.FgLightRed.Println(r.URL)
	color.Green.Println("Whoever receives the note will need the following key to open it:")
	color.FgLightRed.Println(encryptedNote.Key)
}
