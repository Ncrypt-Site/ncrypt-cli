package cypher

import "testing"

func TestEncryptMessage(t *testing.T) {
	note := []byte("It does not do well to dwell on dreams and forget to live")
	key := []byte("3Go2iV@Rb$y*OdIMP3anI89@0i%o!e%w")

	_, err := EncryptNote(note, key)
	if err != nil {
		t.Fatal(err)
	}

	invalidKey := []byte("I'm an invalid key, ha ha ...")
	_, err = EncryptNote(note, invalidKey)
	if err == nil {
		t.Fail()
	}
}
