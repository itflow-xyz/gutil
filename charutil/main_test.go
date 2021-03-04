package charutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const lettersLowercase = "abcdefghijklmnopqrstuvwxyz"
const lettersUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = `~!@#$%^&*()-_=+[{]}\|;:'",<.>/? `

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
	testViaString(t, true, IsNumberOrLetter, lettersUppercase)
	testViaString(t, true, IsNumberOrLetter, lettersLowercase)
	testViaString(t, true, IsNumberOrLetter, numbers)

	testViaString(t, false, IsNumberOrLetter, specialChars)
}

func TestIsNumber(t *testing.T) {
	testViaString(t, true, IsNumber, numbers)

	testViaString(t, false, IsNumber, lettersUppercase)
	testViaString(t, false, IsNumber, lettersLowercase)
	testViaString(t, false, IsNumber, specialChars)
}

func TestIsLetter(t *testing.T) {
	testViaString(t, true, IsLetter, lettersUppercase)
	testViaString(t, true, IsLetter, lettersLowercase)

	testViaString(t, false, IsLetter, numbers)
	testViaString(t, false, IsLetter, specialChars)
}

func TestIsUppercaseLetter(t *testing.T) {
	testViaString(t, true, IsUppercaseLetter, lettersUppercase)

	testViaString(t, false, IsUppercaseLetter, lettersLowercase)
	testViaString(t, false, IsUppercaseLetter, numbers)
	testViaString(t, false, IsUppercaseLetter, specialChars)
}

func TestIsLowercaseLetter(t *testing.T) {
	testViaString(t, true, IsLowercaseLetter, lettersLowercase)

	testViaString(t, false, IsLowercaseLetter, lettersUppercase)
	testViaString(t, false, IsLowercaseLetter, numbers)
	testViaString(t, false, IsLowercaseLetter, specialChars)
}
