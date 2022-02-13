package main

import (
	"strings"
	"testing"
)

func TestShift(t *testing.T) {
	// abcdefghijklmnopqrstuvwxyz
	testString := "uijt jt b uftu."
	expectedString := "this is a test."

	decryptedValue := strings.Map(func(r rune) rune {
		return shift(r, 1)
	}, testString)

	if decryptedValue != expectedString {
		t.Errorf("shift expected %s; got %s", expectedString, decryptedValue)
	}
}
