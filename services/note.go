package services

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"ncrypt-cli/cypher"
	"ncrypt-cli/helpers"
	"ncrypt-cli/models"
)

func CreateNote(note []byte) (models.EncryptedNote, error) {
	if len(note) == 0 {
		return models.EncryptedNote{}, errors.New("note can not be empty")
	}

	key, err := helpers.GenerateRandomString(32)
	iv, err := helpers.GenerateRandomString(16)
	if err != nil {
		return models.EncryptedNote{}, err
	}

	encryptedNote, err := cypher.EncryptNote(note, []byte(key), []byte(iv))
	if err != nil {
		return models.EncryptedNote{}, err
	}
	encryptedNoteB64 := base64.StdEncoding.EncodeToString(encryptedNote)

	noteSha2 := CalculateSHA256(note)

	e := models.EncryptedNote{
		Note: encryptedNoteB64 + "," + noteSha2 + "," + base64.StdEncoding.EncodeToString([]byte(iv)),
		Key:  base64.StdEncoding.EncodeToString([]byte(key)),
	}
	e.Signature = CalculateSHA256([]byte(e.Note))

	return e, nil
}

func CalculateSHA256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
