package excel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIndexToColumn test the IndexToColumn function
func TestIndexToColumn(t *testing.T) {
	assert.Empty(t, IndexToColumn(-1))
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
	assert.Equal(t, -1, ColumnToIndex("1"))
	assert.Equal(t, -1, ColumnToIndex("-"))
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

// TestIndexToRow tests that IndexToRow function
// works correctly
func TestIndexToRow(t *testing.T) {
	assert.Equal(t, 1, IndexToRow(0))
	assert.Equal(t, 2, IndexToRow(1))
}

// TestRowToIndex tests that RowToIndex function
// works correctly
func TestRowToIndex(t *testing.T) {
	assert.Equal(t, 0, RowToIndex(1))
	assert.Equal(t, 1, RowToIndex(2))
}

// TestIndexToAxis with error cases and sample cases
func TestIndexToAxis(t *testing.T) {
	// test error cases
	assert.Empty(t, IndexToAxis(-1, 0))
	assert.Empty(t, IndexToAxis(0, -1))
	assert.Empty(t, IndexToAxis(-1, -1))

	// test sample cases
	assert.Equal(t, "A1", IndexToAxis(0, 0))
	assert.Equal(t, "B2", IndexToAxis(1, 1))
	assert.Equal(t, "Z26", IndexToAxis(25, 25))
	assert.Equal(t, "AA27", IndexToAxis(26, 26))
}

func TestAxisToIndex(t *testing.T) {
	x, y := AxisToIndex("A")
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	x, y = AxisToIndex("1")
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	x, y = AxisToIndex("1A")
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	x, y = AxisToIndex("A1")
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)

	x, y = AxisToIndex("Z26")
	assert.Equal(t, 25, x)
	assert.Equal(t, 25, y)

	x, y = AxisToIndex("AA27")
	assert.Equal(t, 26, x)
	assert.Equal(t, 26, y)
}
