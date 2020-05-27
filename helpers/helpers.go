package helpers

import "errors"

var acceptableSelfDestructValues = map[string]int{
	"1h":  1,
	"3h":  3,
	"6h":  6,
	"12h": 12,
	"1d":  24,
	"2d":  48,
	"3d":  72,
	"7d":  168,
	"1m":  0,
}

func ConvertSelfDestructToInt(i string) (int, error) {
	if v, ok := acceptableSelfDestructValues[i]; ok {
		return v, nil
	}
	return 0, errors.New("invalid input")
}

func ConvertDestructAfterOpeningToBool(i string) (bool, error) {
	if i == "yes" {
		return true, nil
	} else if i == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid input, you should either choose 'yes' or 'no'")
	}
}
