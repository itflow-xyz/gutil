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

func TestRemoveSliceDuplicates(t *testing.T) {
	s := []string{"aaa", "bbb", "ccc"}
	assert.Equal(t, s, RemoveSliceDuplicates(s))

	s = []string{"aaa", "bbb", "aaa"}
	e := []string{"aaa", "bbb"}
	assert.Equal(t, e, RemoveSliceDuplicates(s))
}
