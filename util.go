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

// InRange ...
func InRange(v, min, max int) (c int) {
	if v >= max {
		c = max
		return
	}
	if v <= min {
		c = min
		return
	}
	return v
}
