package excel

import (
	"fmt"

	"github.com/bomaidea/gutil/strutil"
)

// IndexToColumn converts an index (starting at 0)
// to a excel column:
// 0 -> A
// 1 -> B
// 26 -> AA
// 27 -> AB
// 52 -> BA
func IndexToColumn(i int) string {
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
// his index
func ColumnToIndex(c string) int {
	// reverse the column name (the smaller unit to
	// be the first)
	c = strutil.Reverse(c)
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
