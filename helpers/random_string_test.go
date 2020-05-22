package helpers

import (
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	s := GenerateRandomPassword(32)
	if len(s) != 32 {
		t.Fail()
	}
}
