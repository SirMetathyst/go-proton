package configuration

// Provider ...
type Provider func(*C) error

// C ...
type C struct {
	value map[string]interface{}
}

// NewConfiguration ...
func NewConfiguration() *C {
	return &C{make(map[string]interface{}, 0)}
}

var (
	configuration = NewConfiguration()
)

// Singleton ...
func Singleton() *C {
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

// Value ...
func (c *C) Value(key string) interface{} {
	return c.value[key]
}

// Value ...
func Value(key string) interface{} {
	return configuration.Value(key)
}
