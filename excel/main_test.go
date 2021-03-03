package excel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIndexToColumn test the IndexToColumn function
func TestIndexToColumn(t *testing.T) {
	// test first 26 indexes
	for i := 0; i < 26; i++ {
		c := fmt.Sprintf("%c", i+'A')
		assert.Equal(t, c, IndexToColumn(i))
	}

	// test second 26 indexes should have an A in front
	for i := 0; i < 26; i++ {
		c := fmt.Sprintf("A%c", i+'A')
		assert.Equal(t, c, IndexToColumn(i+26))
	}

	// test third 26 indexes should have an A in front
	for i := 0; i < 26; i++ {
		c := fmt.Sprintf("B%c", i+'A')
		assert.Equal(t, c, IndexToColumn(i+52))
	}
}
