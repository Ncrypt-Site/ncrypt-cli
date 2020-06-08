package cypher

import (
	"crypto/aes"
	"crypto/cipher"
)

//EncryptNote encrypt a note with AES256
func EncryptNote(note, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encryptedNote := make([]byte, len(note))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encryptedNote, note)

	return encryptedNote, nil
}
