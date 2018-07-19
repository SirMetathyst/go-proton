package configuration

// KeyBool ...
type KeyBool struct {
	Key   string
	Value *bool
}

// SetBoolP ...
func (c *C) SetBoolP(key string, value *bool) {
	c.SetValue(key, value)
}

// SetBoolP ...
func SetBoolP(key string, value *bool) {
	configuration.SetValue(key, value)
}

// SetBool ...
func (c *C) SetBool(key string, value bool) {
	c.SetValue(key, &value)
}

// SetBool ...
func SetBool(key string, value bool) {
	configuration.SetValue(key, &value)
}

// BoolP ...
func (c *C) BoolP(key string) *bool {
	return c.Value(key).(*bool)
}

// BoolP ...
func BoolP(key string) *bool {
	return configuration.Value(key).(*bool)
}

// AllBool ...
func (c *C) AllBool() []KeyBool {
	slice := make([]KeyBool, 0)
	for k, v := range c.value {
		if bv, ok := v.(*bool); ok {
			slice = append(slice, KeyBool{k, bv})
		}
	}
	return slice
}

// AllBool ...
func AllBool() (Slice []KeyBool) {
	return configuration.AllBool()
}
