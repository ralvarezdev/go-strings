package separator

import (
	"strings"
)

type (
	// Separator represents a string separator.
	Separator string

	// Content represents a pair of left and right separators for content.
	Content struct {
		left  Separator // left is the left separator
		right Separator // right is the right separator
	}

	// Multiline represents separators for formatting multiline content.
	Multiline struct {
		singleElement    Separator // singleElement is the separator for a single element
		multipleElements Separator // multipleElements is the separator for multiple elements
		tabSize          int       // tabSize is the number of tabs to use for indentation
	}
)

// NewContent creates a new Content separator with the specified left and right separators.
//
// Parameters:
//
//	left  - the left separator
//	right - the right separator
//
// Returns:
//
//	A pointer to a Content struct.
func NewContent(left, right Separator) *Content {
	return &Content{
		left,
		right,
	}
}

// Left returns the left separator of the Content.
func (c Content) Left() Separator {
	return c.left
}

// LeftStr returns the left separator as a string.
func (c Content) LeftStr() string {
	return string(c.left)
}

// Right returns the right separator of the Content.
func (c Content) Right() Separator {
	return c.right
}

// RightStr returns the right separator as a string.
func (c Content) RightStr() string {
	return string(c.right)
}

// NewRepeatedContent creates a new Content separator with the same separator for both left and right.
//
// Parameters:
//
//	separator - the separator to use for both left and right
//
// Returns:
//
//	A pointer to a Content struct.
func NewRepeatedContent(separator Separator) *Content {
	return NewContent(separator, separator)
}

// NewMultiline creates a new Multiline separator with the specified single line, multiple line separators, and tab size.
//
// Parameters:
//
//	singleLine - the separator for a single element
//	line       - the separator for multiple elements
//	tabSize    - the number of tabs to use for indentation
//
// Returns:
//
//	A pointer to a Multiline struct.
func NewMultiline(
	singleLine, line Separator,
	tabSize int,
) *Multiline {
	return &Multiline{
		singleLine,
		line,
		tabSize,
	}
}

// SingleElement returns the separator for a single element.
func (m Multiline) SingleElement() Separator {
	return m.singleElement
}

// SingleElementStr returns the single element separator as a string.
func (m Multiline) SingleElementStr() string {
	return string(m.singleElement)
}

// MultipleElements returns the separator for multiple elements.
func (m Multiline) MultipleElements() Separator {
	return m.multipleElements
}

// MultipleElementsStr returns the multiple elements separator as a string.
func (m Multiline) MultipleElementsStr() string {
	return string(m.multipleElements)
}

// TabSize returns the tab size used for indentation.
func (m Multiline) TabSize() int {
	return m.tabSize
}

// TabStr returns the tab separator as a string, repeated tabSize times.
func (m Multiline) TabStr() string {
	return strings.Repeat(string(Tab), m.tabSize)
}
