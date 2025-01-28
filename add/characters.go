package add

import (
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
	"strings"
)

type (
	// CharactersFn is a function that adds characters to a string
	CharactersFn func(*gostringsseparator.Content, string) string
)

// Characters adds some characters to a string
func Characters(
	contentSeparator *gostringsseparator.Content,
	content, leftCharacters, rightCharacters string,
) string {
	if contentSeparator == nil {
		return strings.Join(
			[]string{leftCharacters, content, rightCharacters},
			"",
		)
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

// Brackets adds brackets to a string
func Brackets(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "[", "]")
}

// CurlyBrackets adds curly brackets to a string
func CurlyBrackets(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "{", "}")
}

// Parentheses adds parentheses to a string
func Parentheses(
	contentSeparator *gostringsseparator.Content,
	content string,
) string {
	return Characters(contentSeparator, content, "(", ")")
}
