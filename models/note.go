package models

type Note struct {
	Note                 []byte
	SelfDestruct         int
	DestructAfterOpening bool
}

type EncryptedNote struct {
	Note string
	Key  string
}
