package add

import (
	"strings"

	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

type (
	// CharactersFn is a function type that adds characters to a string.
	//
	// Parameters:
	//
	//   contentSeparator - pointer to a Content struct for separators
	//   content          - the string to which characters will be added
	//
	// Returns:
	//
	//   The resulting string after adding characters.
	CharactersFn func(
		contentSeparator *gostringsseparator.Content,
		content string,
	) string
)

// Characters adds left and right characters to the content string, optionally using a separator.
//
// Parameters:
//
//	contentSeparator - pointer to a Content struct for separators (can be nil)
//	content          - the main string to wrap
//	leftCharacters   - characters to add to the left of the content
//	rightCharacters  - characters to add to the right of the content
//
// Returns:
//
//	The resulting string with added characters and optional separators.
func Characters(
	contentSeparator *gostringsseparator.Content,
	content, leftCharacters, rightCharacters string,
) string {
	if contentSeparator == nil {
		return leftCharacters + content + rightCharacters
	}

	return strings.Join(
		[]string{
			leftCharacters,
			contentSeparator.LeftStr(),
			content,
			contentSeparator.RightStr(),
			rightCharacters,
		}, "",
	)
}

// Brackets adds square brackets around the content string, optionally using a separator.
//
// Parameters:
//
//	contentSeparator - pointer to a Content struct for separators (can be nil)
//	content          - the string to wrap in brackets
//
// Returns:
//
//	The resulting string wrapped in square brackets.
func Brackets(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "[", "]")
}

// CurlyBrackets adds curly brackets around the content string, optionally using a separator.
//
// Parameters:
//
//	contentSeparator - pointer to a Content struct for separators (can be nil)
//	content          - the string to wrap in curly brackets
//
// Returns:
//
//	The resulting string wrapped in curly brackets.
func CurlyBrackets(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "{", "}")
}

// Parentheses adds parentheses around the content string, optionally using a separator.
//
// Parameters:
//
//	contentSeparator - pointer to a Content struct for separators (can be nil)
//	content          - the string to wrap in parentheses
//
// Returns:
//
//	The resulting string wrapped in parentheses.
func Parentheses(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "(", ")")
}
