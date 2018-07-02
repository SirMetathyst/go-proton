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

// GetStringSliceP ...
func (c *C) GetStringSliceP(key string) *[]string {
	return c.GetValue(key).(*[]string)
}

// GetStringSliceP ...
func GetStringSliceP(key string) *[]string {
	return configuration.GetValue(key).(*[]string)
}

// GetAllStringSlice ...
func (c *C) GetAllStringSlice() []KeyStringSlice {
	slice := make([]KeyStringSlice, 0)
	for k, v := range c.value {
		if ssv, ok := v.(*[]string); ok {
			slice = append(slice, KeyStringSlice{k, ssv})
		}
	}
	return slice
}

// GetAllStringSlice ...
func GetAllStringSlice() []KeyStringSlice {
	return configuration.GetAllStringSlice()
}
