package proton

import "unicode"

// containsWhitespace ...
func containsWhitespace(v string) bool {
	for _, vv := range v {
		if unicode.IsSpace(vv) {
			return true
		}
	}
	return false
}
