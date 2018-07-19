package configurationprovider

import (
	"os"
	"strings"

	"github.com/SirMetathyst/proton/configuration"
)

// Env ...
func Env(c *configuration.C) error {
	for _, b := range c.AllBool() {
		*b.Value = getBool(b.Key, *b.Value)
	}
	return nil
}

// getBool ...
func getBool(Key string, DefaultValue bool) bool {
	Value := os.Getenv(Key)
	if Value == "" {
		return DefaultValue
	}
	if strings.EqualFold(Value, "true") {
		return true
	}
	return false
}
