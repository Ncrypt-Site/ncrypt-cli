package helpers

import (
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	r1, err := GenerateRandomString(32)
	if err != nil {
		t.Fatal(err)
	}

	if len(r1) != 32 {
		t.Fail()
	}

	r2, err := GenerateRandomString(32)
	if r1 == r2 {
		t.Fail()
	}
}
