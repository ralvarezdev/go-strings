package go_strings

import (
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

// Repeat repeats a string n times
func Repeat(s string, n int, separator gostringsseparator.Separator) string {
	var result string
	separatorStr := string(separator)
	for i := 0; i < n; i++ {
		result += s + separatorStr
	}
	return result
}

// RepeatAsArray repeats a string n times and returns an array of strings
func RepeatAsArray(s string, n int) *[]string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, s)
	}
	return &result
}
