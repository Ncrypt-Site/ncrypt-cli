package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptNote(note []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedText := gcm.Seal(nonce, nonce, note, nil)
	return encryptedText, nil
}
