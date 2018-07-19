package configuration

// KeyStringSlice ...
type KeyStringSlice struct {
	Key   string
	Value *[]string
}

// SetStringSliceP ...
func (c *C) SetStringSliceP(key string, value *[]string) {
	c.SetValue(key, value)
}

// SetStringSliceP ...
func SetStringSliceP(key string, value *[]string) {
	configuration.SetValue(key, value)
}

// SetStringSlice ...
func (c *C) SetStringSlice(key string, value []string) {
	c.SetValue(key, &value)
}

// SetStringSlice ...
func SetStringSlice(key string, value []string) {
	configuration.SetValue(key, &value)
}

// StringSliceP ...
func (c *C) StringSliceP(key string) *[]string {
	return c.Value(key).(*[]string)
}

// StringSliceP ...
func StringSliceP(key string) *[]string {
	return configuration.Value(key).(*[]string)
}

// AllStringSlice ...
func (c *C) AllStringSlice() []KeyStringSlice {
	slice := make([]KeyStringSlice, 0)
	for k, v := range c.value {
		if ssv, ok := v.(*[]string); ok {
			slice = append(slice, KeyStringSlice{k, ssv})
		}
	}
	return slice
}

// AllStringSlice ...
func AllStringSlice() []KeyStringSlice {
	return configuration.AllStringSlice()
}
