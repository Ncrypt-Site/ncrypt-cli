package services

import (
	"crypto/sha256"
	"encoding/base64"
	"reflect"
	"testing"
)

func TestGenerateSHA256(t *testing.T) {
	message := []byte("sample message")

	shaB64 := GenerateSHA256(message)
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
