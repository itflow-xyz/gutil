package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRgbToHex(t *testing.T) {
	assert.Equal(t, "#000000", RgbToHex(0, 0, 0))
	assert.Equal(t, "#0F0F0F", RgbToHex(15, 15, 15))
	assert.Equal(t, "#FFFFFF", RgbToHex(255, 255, 255))
}
