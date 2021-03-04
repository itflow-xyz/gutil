package charutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testViaString is a test helper, it test the function
// given as parameter [f] with all the runes of the
// string [s], if the bool [e] is expected true from the
// function otherways false
func testViaString(t *testing.T, e bool, f func(rune) bool, s string) {
	for _, v := range s {
		if e {
			assert.True(t, f(v))
		} else {
			assert.False(t, f(v))
		}
	}
}

func TestIsNumberOrLetter(t *testing.T) {
	testViaString(t, true, IsNumberOrLetter, LettersUppercase)
	testViaString(t, true, IsNumberOrLetter, LettersLowercase)
	testViaString(t, true, IsNumberOrLetter, Numbers)

	testViaString(t, false, IsNumberOrLetter, SpecialChars)
}

func TestIsNumber(t *testing.T) {
	testViaString(t, true, IsNumber, Numbers)

	testViaString(t, false, IsNumber, LettersUppercase)
	testViaString(t, false, IsNumber, LettersLowercase)
	testViaString(t, false, IsNumber, SpecialChars)
}

func TestIsLetter(t *testing.T) {
	testViaString(t, true, IsLetter, LettersUppercase)
	testViaString(t, true, IsLetter, LettersLowercase)

	testViaString(t, false, IsLetter, Numbers)
	testViaString(t, false, IsLetter, SpecialChars)
}

func TestIsUppercaseLetter(t *testing.T) {
	testViaString(t, true, IsUppercaseLetter, LettersUppercase)

	testViaString(t, false, IsUppercaseLetter, LettersLowercase)
	testViaString(t, false, IsUppercaseLetter, Numbers)
	testViaString(t, false, IsUppercaseLetter, SpecialChars)
}

func TestIsLowercaseLetter(t *testing.T) {
	testViaString(t, true, IsLowercaseLetter, LettersLowercase)

	testViaString(t, false, IsLowercaseLetter, LettersUppercase)
	testViaString(t, false, IsLowercaseLetter, Numbers)
	testViaString(t, false, IsLowercaseLetter, SpecialChars)
}
