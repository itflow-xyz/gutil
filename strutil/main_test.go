package strutil

import (
	"testing"

	"github.com/bomaidea/gutil/charutil"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "AA", Reverse("AA"))
	assert.Equal(t, "CBA", Reverse("ABC"))
	assert.Equal(t, "ABA", Reverse("ABA"))
}

func TestContainsOnlyLetter(t *testing.T) {
	assert.True(t, ContainsOnlyLetter(charutil.LettersUppercase))
	assert.True(t, ContainsOnlyLetter(charutil.LettersLowercase))

	assert.False(t, ContainsOnlyLetter(charutil.Numbers))
	assert.False(t, ContainsOnlyLetter(charutil.SpecialChars))
}
