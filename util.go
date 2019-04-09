package proton

import "unicode"

// ContainsWhitespace ...
func ContainsWhitespace(v string) bool {
	for _, vv := range v {
		if unicode.IsSpace(vv) {
			return true
		}
	}
	return false
}
