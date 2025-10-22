package gostrings

import (
	"strings"
)

// ToCamelCase converts a string to CamelCase.
//
// Parameters:
//
//   - s: The string to convert.
//
// Returns:
//
//   - string: The converted string.
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}
	parts := strings.FieldsFunc(
		s, func(r rune) bool {
			return r == '_' || r == '-' || r == ' '
		},
	)
	for i, part := range parts {
		if part != "" {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

// ToSnakeCase converts a string to snake_case.
//
// Parameters:
//
//   - s: The string to convert.
//
// Returns:
//
//   - string: The converted string.
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' && (s[i-1] >= 'a' && s[i-1] <= 'z') {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
