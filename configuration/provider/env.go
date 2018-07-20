package configurationprovider

import (
	"os"
	"strings"

	"github.com/SirMetathyst/go-blackboard"
)

// Env ...
func Env(bb *blackboard.BB) error {
	for _, b := range bb.AllBool() {
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
