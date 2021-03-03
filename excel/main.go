package excel

import (
	"fmt"
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
