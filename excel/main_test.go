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

// TestColumnToIndex test important values for check
// that ColumnToIndex function works correctly
func TestColumnToIndex(t *testing.T) {
	assert.Equal(t, 0, ColumnToIndex("A"))
	assert.Equal(t, 1, ColumnToIndex("B"))
	assert.Equal(t, 25, ColumnToIndex("Z"))
	assert.Equal(t, 26, ColumnToIndex("AA"))
	assert.Equal(t, 27, ColumnToIndex("AB"))
	assert.Equal(t, 52, ColumnToIndex("BA"))
	assert.Equal(t, 701, ColumnToIndex("ZZ"))
	assert.Equal(t, 702, ColumnToIndex("AAA"))
	assert.Equal(t, 703, ColumnToIndex("AAB"))
}
