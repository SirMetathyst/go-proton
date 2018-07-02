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

// GetStringP ...
func (c *C) GetStringP(key string) *string {
	return c.GetValue(key).(*string)
}

// GetStringP ...
func GetStringP(key string) *string {
	return configuration.GetValue(key).(*string)
}

// GetAllString ...
func (c *C) GetAllString() []KeyString {
	slice := make([]KeyString, 0)
	for k, v := range c.value {
		if sv, ok := v.(*string); ok {
			slice = append(slice, KeyString{k, sv})
		}
	}
	return slice
}

// GetAllString ...
func GetAllString() (Slice []KeyString) {
	return configuration.GetAllString()
}
