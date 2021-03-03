package strutil

// Reverse reverts a string ASD -> DSA, abc -> cba
func Reverse(s string) string {
	var out string

	for _, v := range s {
		out = string(v) + out
	}

	return out
}
