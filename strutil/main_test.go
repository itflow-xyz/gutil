package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "AA", Reverse("AA"))
	assert.Equal(t, "CBA", Reverse("ABC"))
	assert.Equal(t, "ABA", Reverse("ABA"))
}
