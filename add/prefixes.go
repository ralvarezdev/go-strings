package add

import (
	"strings"

	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

// Prefixes returns a string by joining the given prefixes and content using the specified separator.
//
// Parameters:
//
//	content   - the main string to append after the prefixes
//	separator - the separator to use between elements
//	prefixes  - variadic list of prefix strings to prepend to the content
//
// Returns:
//
//	A single string with all prefixes and the content joined by the separator.
func Prefixes(
	content string,
	separator gostringsseparator.Separator,
	prefixes ...string,
) string {
	// Push content to the end of the prefixes string slice
	prefixes = append(prefixes, content)

	return strings.Join(prefixes, string(separator))
}
