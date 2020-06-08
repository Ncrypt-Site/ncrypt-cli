package services

import (
	"encoding/json"
	"errors"
	"ncrypt-cli/models"
)

//CreateNote create a note and send to API
func CreateNote(h HttpService, n models.Note) (models.NoteCreatedResponse, string, error) {
	encryptedNote, err := createSecureNote(n.Note)
	if err != nil {
		return models.NoteCreatedResponse{}, "", errors.New("an error occurred, please try again")
	}

	model := models.SecureMessageRequest{
		Note:                 encryptedNote.Note,
		SelfDestruct:         n.SelfDestruct,
		DestructAfterOpening: n.DestructAfterOpening,
	}

	resp, err := h.SendRequest(model)
	if err != nil {
		return models.NoteCreatedResponse{}, "", err
	}

	r := models.NoteCreatedResponse{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return models.NoteCreatedResponse{}, "", err
	}

	return r, encryptedNote.Key, nil
}
