package helpers

import (
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	s, err := GenerateRandomString(32)
	if err != nil {
		t.Fatal(err)
	}

	if len(s) != 32 {
		t.Fail()
	}
}
