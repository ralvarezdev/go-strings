package format

import (
	"strings"

	gostringsadd "github.com/ralvarezdev/go-strings/add"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
)

// StringArray returns a formatted string containing all elements of the given string array.
//
// Parameters:
//
//	multilineSeparator - pointer to a Multiline separator that defines how to format the string
//	stringArray        - pointer to a slice of strings to be formatted
//	addCharactersFn    - function that adds characters to the formatted string
//
// Returns:
//
//	A formatted string with all the strings in the array, or an empty string if the array is nil or empty.
func StringArray(
	multilineSeparator *gostringsseparator.Multiline,
	stringArray *[]string,
	addCharactersFn gostringsadd.CharactersFn,
) string {
	// Check if the stringArray is nil or empty, or the addCharactersFn is nil
	if stringArray == nil || len(*stringArray) == 0 || addCharactersFn == nil {
		return ""
	}

	// Check if there is only one element
	if len(*stringArray) == 1 {
		return addCharactersFn(
			gostringsseparator.NewRepeatedContent(multilineSeparator.SingleElement()),
			(*stringArray)[0],
		)
	} else {
		var formattedDetails strings.Builder

		// Separators
		elementSeparatorStr := multilineSeparator.MultipleElementsStr()
		tabSeparatorStr := multilineSeparator.TabStr()
		elementAndTabSeparatorStr := elementSeparatorStr + tabSeparatorStr

		// Add formatted details
		formattedDetails.WriteString(tabSeparatorStr)
		for i, str := range *stringArray {
			formattedDetails.WriteString(str)

			if i < len(*stringArray)-1 {
				formattedDetails.WriteString(elementAndTabSeparatorStr)
			}
		}

		return addCharactersFn(
			gostringsseparator.NewRepeatedContent(multilineSeparator.MultipleElements()),
			formattedDetails.String(),
		)
	}
}

// ErrorArray returns a formatted string containing all error messages from the given error array.
//
// Parameters:
//
//	multilineSeparator - pointer to a Multiline separator that defines how to format the string
//	errorArray         - pointer to a slice of errors to be formatted
//	addCharactersFn    - function that adds characters to the formatted string
//
// Returns:
//
//	A formatted string with all the error messages, or an empty string if the array is nil or empty.
func ErrorArray(
	multilineSeparator *gostringsseparator.Multiline,
	errorArray *[]error,
	addCharactersFn gostringsadd.CharactersFn,
) string {
	mappedErrorArray := gostringsconvert.ErrorArrayToStringArray(errorArray)
	return StringArray(
		multilineSeparator,
		mappedErrorArray,
		addCharactersFn,
	)
}
