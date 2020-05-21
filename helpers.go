package main

import "errors"

var acceptableSelfDestructValues = map[string]int{
	"1 hour":   1,
	"3 hours":  3,
	"6 hours":  6,
	"12 hours": 12,
	"1 day":    24,
	"2 days":   48,
	"3 days":   72,
	"7 days":   168,
	"1 month":  0,
}

func convertSelfDestructToInt(i string) (int, error) {
	if v, ok := acceptableSelfDestructValues[i]; ok {
		return v, nil
	}
	return 0, errors.New("invalid input")
}

func convertDestructAfterOpeningToBool(i string) (bool, error) {
	if i == "yes" {
		return true, nil
	} else if i == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid input, you should either choose 'yes' or 'no'")
	}
}
