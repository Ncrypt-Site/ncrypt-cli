package models

//Note model for http request
type Note struct {
	Note                 []byte
	SelfDestruct         int
	DestructAfterOpening bool
}

//EncryptedNote model for returning data to the user
type EncryptedNote struct {
	Note string
	Key  string
}
