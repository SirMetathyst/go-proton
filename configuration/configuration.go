package configuration

var (
	configuration = New()
)

// Provider ...
type Provider func(*C) error

// C ...
type C struct {
	value map[string]interface{}
}

// New ...
func New() *C {
	return &C{make(map[string]interface{}, 0)}
}

// Get ...
func Get() *C {
	return configuration
}

// SetValue ...
func (c *C) SetValue(key string, value interface{}) {
	c.value[key] = value
}

// SetValue ...
func SetValue(key string, value interface{}) {
	configuration.SetValue(key, value)
}

// GetValue ...
func (c *C) GetValue(key string) interface{} {
	return c.value[key]
}

// GetValue ...
func GetValue(key string) interface{} {
	return configuration.GetValue(key)
}
