package format

import (
	gostringsadd "github.com/ralvarezdev/go-strings/add"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	gostringsseparator "github.com/ralvarezdev/go-strings/separator"
	"strings"
)

// StringArray returns a string with all the strings in the array formatted
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
			gostringsseparator.NewRepeatedContent(multilineSeparator.SingleLine()),
			(*stringArray)[0],
		)
	} else {
		var formattedDetails strings.Builder

		// Separators
		lineSeparatorStr := multilineSeparator.LineStr()
		tabSeparatorStr := multilineSeparator.TabStr()
		lineAndTabSeparatorStr := lineSeparatorStr + tabSeparatorStr

		// Add formatted details
		formattedDetails.WriteString(tabSeparatorStr)
		for i, str := range *stringArray {
			formattedDetails.WriteString(str)

			if i < len(*stringArray)-1 {
				formattedDetails.WriteString(lineAndTabSeparatorStr)
			}
		}

		return addCharactersFn(
			gostringsseparator.NewRepeatedContent(multilineSeparator.Line()),
			formattedDetails.String(),
		)
	}
}

// ErrorArray returns a string with all the errors in the array formatted
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
