package helpers

import "testing"

func TestConvertDestructAfterOpeningToBool(t *testing.T) {
	_, err := ConvertSelfDestructToInt("1h")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ConvertSelfDestructToInt("100 hours")
	if err == nil {
		t.Fail()
	}
}

func TestConvertSelfDestructToInt(t *testing.T) {
	_, err := ConvertDestructAfterOpeningToBool("no")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ConvertDestructAfterOpeningToBool("yes")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ConvertDestructAfterOpeningToBool("oops")
	if err == nil {
		t.Fail()
	}
}
