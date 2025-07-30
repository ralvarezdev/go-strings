package add

import (
	"strings"

	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

// Prefixes returns a string with the given content and prefixes
func Prefixes(
	content string,
	separator gostringsseparator.Separator,
	prefixes ...string,
) string {
	// Push content to the end of the prefixes string slice
	prefixes = append(prefixes, content)

	return strings.Join(prefixes, string(separator))
}
