package charutil

// IsSpecialChar returns true if the rune given as
// parameter is not a special char, it could be in
// between this ranges: A-Z a-z 0-9 or it can be a
// space
func IsSpecialChar(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '0'
}

// IsNumber returns true if the rune given as
// parameter is a number
func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsLetter returns true if the rune given as
// parameter is a letter (could be upper or lower
// case)
func IsLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
