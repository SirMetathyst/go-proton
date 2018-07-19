package model

import "strings"

// String ...
type String string

// String ...
func (s String) String() string {
	return string(s)
}

// EqualTo ...
func (s String) EqualTo(v string) bool {
	return string(s) == v
}

// ToUpperFirst ...
func (s String) ToUpperFirst() String {
	return String(strings.ToUpper(string(s)[:1]) + string(s)[1:])
}

// ToUpper ...
func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

// ToLowerFirst ...
func (s String) ToLowerFirst() String {
	if len(s) > 0 {
		return String(strings.ToLower(string(s)[:1]) + string(s)[1:])
	}
	return s
}

// ToLower ...
func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

// WithContextSuffix ...
func (s String) WithContextSuffix() String {
	if strings.HasSuffix(s.String(), "Context") {
		return s
	}
	return s + String("Context")
}

// WithoutContextSuffix ...
func (s String) WithoutContextSuffix() String {
	if strings.HasSuffix(s.String(), "Context") {
		return String(strings.TrimSuffix(s.String(), "Context"))
	}
	return s
}

// WithComponentSuffix ...
func (s String) WithComponentSuffix() String {
	if strings.HasSuffix(s.String(), "Component") {
		return s
	}
	return s + String("Component")
}

// WithoutComponentSuffix ...
func (s String) WithoutComponentSuffix() String {
	if strings.HasSuffix(s.String(), "Component") {
		return String(strings.TrimSuffix(s.String(), "Component"))
	}
	return s
}
