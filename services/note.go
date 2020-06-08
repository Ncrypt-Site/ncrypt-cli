package services

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"ncrypt-cli/cypher"
	"ncrypt-cli/helpers"
	"ncrypt-cli/models"
)

func createSecureNote(note []byte) (models.EncryptedNote, error) {
	if len(note) == 0 {
		return models.EncryptedNote{}, errors.New("note can not be empty")
	}

	key, err := helpers.GenerateRandomString(32)
	if err != nil {
		return models.EncryptedNote{}, err
	}

	iv, err := helpers.GenerateRandomString(16)
	if err != nil {
		return models.EncryptedNote{}, err
	}

	encryptedNote, err := cypher.EncryptNote(note, []byte(key), []byte(iv))
	if err != nil {
		return models.EncryptedNote{}, err
	}
	encryptedNoteB64 := base64.StdEncoding.EncodeToString(encryptedNote)

	noteSha2 := calculateSHA256(note)

	e := models.EncryptedNote{
		Note: encryptedNoteB64 + "," + noteSha2 + "," + base64.StdEncoding.EncodeToString([]byte(iv)),
		Key:  base64.StdEncoding.EncodeToString([]byte(key)),
	}
	encryptedNoteSignature := calculateSHA256([]byte(e.Note))
	e.Note += "," + encryptedNoteSignature

	return e, nil
}

func calculateSHA256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
