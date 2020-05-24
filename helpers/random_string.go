package helpers

import (
	"math/rand"
)

const CharSets = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	l := len(CharSets)

	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", nil
	}
	s := make([]byte, n)
	for i, v := range b {
		s[i] = CharSets[v%uint8(l)]
	}

	return string(s), nil
}
