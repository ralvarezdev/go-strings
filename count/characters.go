package count

// Numbers counts the number of numbers in a string
//
// Parameters:
//
//	s - the string to count numbers in
//
// Returns:
//
//	The count of numbers in the string
func Numbers(s string) int {
	count := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			count++
		}
	}
	return count
}

// Capital counts the number of capital letters in a string
//
// Parameters:
//
//	s - the string to count capital letters in
//
// Returns:
//
//	The count of capital letters in the string
func Capital(s string) int {
	count := 0
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			count++
		}
	}
	return count
}

// Caps counts the number of capital letters in a string
//
// Parameters:
//
//	s - the string to count capital letters in
//
// Returns:
//
//	The count of capital letters in the string
func Caps(s string) int {
	return Capital(s)
}

// Special counts the number of special characters in a string
//
// Parameters:
//
//	s - the string to count special characters in
//
// Returns:
//
//	The count of special characters in the string
func Special(s string) int {
	count := 0
	for _, r := range s {
		if (r >= '!' && r <= '/') ||
			(r >= ':' && r <= '@') ||
			(r >= '[' && r <= '`') ||
			(r >= '{' && r <= '~') {
			count++
		}
	}
	return count
}

// Lowercase counts the number of lowercase letters in a string
//
// Parameters:
//
//	s - the string to count lowercase letters in
//
// Returns:
//
//	The count of lowercase letters in the string
func Lowercase(s string) int {
	count := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count++
		}
	}
	return count
}

// Alphabetic counts the number of alphabetic characters in a string
//
// Parameters:
//
//	s - the string to count alphabetic characters in
//
// Returns:
//
//	The count of alphabetic characters in the string
func Alphabetic(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}
	return count
}

// Alphanumeric counts the number of alphanumeric characters in a string
//
// Parameters:
//
//	s - the string to count alphanumeric characters in
//
// Returns:
//
//	The count of alphanumeric characters in the string
func Alphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return count
}
