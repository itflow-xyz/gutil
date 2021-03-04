package charutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsSpecialChar tests that the func works correctly
func TestIsSpecialChar(t *testing.T) {
	// test true chars
	s := "ABCXYZabcxyz012789"
	for _, v := range s {
		assert.True(t, IsSpecialChar(v))
	}

	// test special chars
	s = "!@#$%^&*()_+}{:\\\"|?><~`"
	for _, v := range s {
		assert.False(t, IsSpecialChar(v))
	}
}

// TestIsNumber test that the function returns
// the right values with the right paramters
// values
func TestIsNumber(t *testing.T) {
	// test with numbers
	s := "0123456789"
	for _, v := range s {
		assert.True(t, IsNumber(v))
	}

	// test with letters and special chars
	s = "ABCXYZabcxyz!@#$%^&*(()_+}{\\:;[]\"/.,><?"
	for _, v := range s {
		assert.False(t, IsNumber(v))
	}
}

func TestIsLetter(t *testing.T) {
	// test with letters
	s := "ABCXYZabcxyz"
	for _, v := range s {
		assert.True(t, IsLetter(v))
	}

	// test with numbers and special chars
	s = "012789!#$%^&*()_+}{[]:;\"\\|/.,?><`~"
	for _, v := range s {
		assert.False(t, IsLetter(v))
	}
}
