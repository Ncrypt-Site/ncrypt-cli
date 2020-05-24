package services

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateSHA256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
