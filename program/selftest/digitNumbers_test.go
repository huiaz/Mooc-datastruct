package selftest

import "testing"

func TestDigitNumbers(t *testing.T) {
	t.Log(digitNumbers("123456789"))
}

func TestDoubleStringNumber(t *testing.T) {
	t.Log(doubleStringNumber("9999999999999999999"))
}
