package go_strings

import (
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

// Repeat returns a new string consisting of the input string `s` repeated `n` times,
// separated by the specified separator.
//
// Parameters:
//
//	s         - the string to repeat
//	n         - the number of times to repeat the string
//	separator - the separator to insert between repetitions
//
// Returns:
//
//	A single string with `s` repeated `n` times, separated by `separator`.
func Repeat(s string, n int, separator gostringsseparator.Separator) string {
	var result string
	separatorStr := string(separator)
	for i := 0; i < n; i++ {
		result += s + separatorStr
	}
	return result
}

// RepeatAsArray returns a pointer to a slice containing the input string `s` repeated `n` times.
//
// Parameters:
//
//	s - the string to repeat
//	n - the number of times to repeat the string
//
// Returns:
//
//	A pointer to a slice of strings, each element being `s`.
func RepeatAsArray(s string, n int) *[]string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, s)
	}
	return &result
}
