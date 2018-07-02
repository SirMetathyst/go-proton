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

// GetBoolP ...
func (c *C) GetBoolP(key string) *bool {
	return c.GetValue(key).(*bool)
}

// GetBoolP ...
func GetBoolP(key string) *bool {
	return configuration.GetValue(key).(*bool)
}

// GetAllBool ...
func (c *C) GetAllBool() []KeyBool {
	slice := make([]KeyBool, 0)
	for k, v := range c.value {
		if bv, ok := v.(*bool); ok {
			slice = append(slice, KeyBool{k, bv})
		}
	}
	return slice
}

// GetAllBool ...
func GetAllBool() (Slice []KeyBool) {
	return configuration.GetAllBool()
}
