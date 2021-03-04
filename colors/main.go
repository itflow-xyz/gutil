package colors

import (
	"fmt"
)

// RgbToHex converts rgb color values to HEX
// colors
func RgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
