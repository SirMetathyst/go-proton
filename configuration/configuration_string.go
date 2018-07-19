package configuration

// KeyString ...
type KeyString struct {
	Key   string
	Value *string
}

// SetStringP ...
func (c *C) SetStringP(key string, value *string) {
	c.SetValue(key, value)
}

// SetStringP ...
func SetStringP(key string, value *string) {
	configuration.SetValue(key, value)
}

// SetString ...
func (c *C) SetString(key string, value string) {
	c.SetValue(key, &value)
}

// SetString ...
func SetString(key string, value string) {
	configuration.SetValue(key, &value)
}

// StringP ...
func (c *C) StringP(key string) *string {
	return c.Value(key).(*string)
}

// StringP ...
func StringP(key string) *string {
	return configuration.Value(key).(*string)
}

// AllString ...
func (c *C) AllString() []KeyString {
	slice := make([]KeyString, 0)
	for k, v := range c.value {
		if sv, ok := v.(*string); ok {
			slice = append(slice, KeyString{k, sv})
		}
	}
	return slice
}

// AllString ...
func AllString() (Slice []KeyString) {
	return configuration.AllString()
}
