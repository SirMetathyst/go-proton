package proton

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
