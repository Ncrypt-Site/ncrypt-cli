package models

type EncryptedNote struct {
	Note      string
	Key       string
	Signature string
}
