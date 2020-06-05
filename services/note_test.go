package services

import (
	"crypto/sha256"
	"encoding/base64"
	"reflect"
	"testing"
)

func TestGenerateSHA256(t *testing.T) {
	message := []byte("Ever have that feeling where you’re not sure if you’re awake or dreaming?")

	shaB64 := CalculateSHA256(message)
	sha, err := base64.StdEncoding.DecodeString(shaB64)
	if err != nil {
		t.Fatal(err)
	}

	h := sha256.New()
	h.Write(message)

	if !reflect.DeepEqual(h.Sum(nil), sha) {
		t.Fail()
	}
}

// todo: this test needs improvement.
func TestCreateSecureNote(t *testing.T) {
	message := []byte("Throughout human history, we have been dependent on machines to survive." +
		" Fate, it seems, is not without a sense of irony.")
	_, err := createSecureNote(message)
	if err != nil {
		t.Fatal(err)
	}

	_, err = createSecureNote([]byte(""))
	if err == nil {
		t.Fail()
	}
}
