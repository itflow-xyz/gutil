package charutil

// LettersLowercase list of all letters lowercase,
// alphabeth order
const LettersLowercase = "abcdefghijklmnopqrstuvwxyz"

// LettersUppercase list of all letters uppercase,
// alphabeth order
const LettersUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Numbers list of all numbers
const Numbers = "0123456789"

// SpecialChars list of most used special chars
const SpecialChars = `~!@#$%^&*()-_=+[{]}\|;:'",<.>/? `

// Letters is the LettersUppercase concatenated to
// LettersLowercase
const Letters = LettersUppercase + LettersLowercase

// IsNumberOrLetter returns true if the rune given as
// parameter is a letter (upper or lower case) or a
// a number, it could be in between this ranges: A-Z
// a-z 0-9 or it can be a space
func IsNumberOrLetter(r rune) bool {
	return IsNumber(r) || IsLetter(r)
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
	return IsUppercaseLetter(r) || IsLowercaseLetter(r)
}

// IsUppercaseLetter returns true if the rune
// given as parameter is a upper case letter
func IsUppercaseLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsLowercaseLetter returns true if the rune
// given as parameter is lower case letter
func IsLowercaseLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}
