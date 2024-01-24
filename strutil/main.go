package strutil

import (
	"github.com/itflow-xyz/gutil/charutil"
)

// Reverse reverts a string ASD -> DSA, abc -> cba
func Reverse(s string) string {
	var out string

	for _, v := range s {
		out = string(v) + out
	}

	return out
}

// ContainsOnlyLetter returns true if the string
// given as parameter contains only strings
func ContainsOnlyLetter(s string) bool {
	for _, v := range s {
		if !charutil.IsLetter(v) {
			return false
		}
	}
	return true
}

func RemoveSliceDuplicates(s []string) []string {
	check := make(map[string]int)
	res := make([]string, 0)

	for _, v := range s {
		check[v] = 1
	}

	for l, _ := range check {
		res = append(res, l)
	}

	return res
}
