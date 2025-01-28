package separator

import (
	"strings"
)

type (
	// Separator is the separator for a string
	Separator string

	// Content is the separator for the content
	Content struct {
		left  Separator
		right Separator
	}

	// Multiline is the separator for the multiline content
	Multiline struct {
		singleLine Separator
		line       Separator
		tabSize    int
	}
)

// Separator constants
const (
	Space     Separator = " "
	Comma     Separator = ","
	NewLine   Separator = "\n"
	Tab       Separator = "\t"
	Dots      Separator = ":"
	Slash     Separator = "/"
	Backslash Separator = "\\"
)

// NewContent creates a new content separator
func NewContent(left, right Separator) *Content {
	return &Content{
		left,
		right,
	}
}

// Left returns the left separator
func (c *Content) Left() Separator {
	return c.left
}

// LeftStr returns the left separator as a string
func (c *Content) LeftStr() string {
	return string(c.left)
}

// Right returns the right separator
func (c *Content) Right() Separator {
	return c.right
}

// RightStr returns the right separator as a string
func (c *Content) RightStr() string {
	return string(c.right)
}

// NewRepeatedContent creates a new content separator with the same separator
func NewRepeatedContent(separator Separator) *Content {
	return NewContent(separator, separator)
}

// NewMultiline creates a new multiline separator
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

// SingleLine returns the single line separator
func (m *Multiline) SingleLine() Separator {
	return m.singleLine
}

// SingleLineStr returns the single line separator as a string
func (m *Multiline) SingleLineStr() string {
	return string(m.singleLine)
}

// Line returns the line separator
func (m *Multiline) Line() Separator {
	return m.line
}

// LineStr returns the line separator as a string
func (m *Multiline) LineStr() string {
	return string(m.line)
}

// TabSize returns the tab size
func (m *Multiline) TabSize() int {
	return m.tabSize
}

// TabStr returns the tab as a string
func (m *Multiline) TabStr() string {
	return strings.Repeat(string(Tab), m.tabSize)
}
