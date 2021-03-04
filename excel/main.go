package excel

import (
	"fmt"
	"strings"

	"github.com/bomaidea/gutil/strutil"
)

// IndexToColumn converts an index (starting at 0)
// to a excel column:
// 0 -> A
// 1 -> B
// 26 -> AA
// 27 -> AB
// 52 -> BA
// if the index is negative returns empty
func IndexToColumn(i int) string {
	// if the index is negative returns empty
	if i < 0 {
		return ""
	}

	s := ""
	for i > 0 || len(s) == 0 {
		y := i / ('Z' - 'A' + 1)
		r := i % ('Z' - 'A' + 1)
		i = y
		if len(s) > 0 {
			r--
		}
		s = fmt.Sprintf("%c", r+'A') + s
	}

	return s
}

// ColumnToIndex converts an excel column name to
// his index, if the name contains special chars
// or numbers returns -1
func ColumnToIndex(c string) int {
	// reverse the column name (the smaller unit to
	// be the first)
	c = strutil.Reverse(c)
	// set to upppercase
	c = strings.ToUpper(c)

	// check if the column name contains only letter
	if !strutil.ContainsOnlyLetter(c) {
		return -1
	}

	// create multiplicator at 1
	m := 1
	var i int
	// for each "letter" of the column
	for j, v := range c {
		// calculate the letter weight with the multiplicator
		if j > 0 {
			i += (1 + int(v-'A')) * m
		} else {
			i += int(v-'A') * m
		}
		// update the multiplicator
		m *= ('Z' - 'A' + 1)
	}
	return i
}

// IndexToRow converts the generical Index numeration
// to excel numeration
// - generical numeration starts at 0
// - excel numeration starts at 1
func IndexToRow(x int) int {
	return x + 1
}

// IndexToRow converts the excel numeration to
// generical Index numeration
// - generical numeration starts at 0
// - excel numeration starts at 1
func RowToIndex(x int) int {
	return x - 1
}
