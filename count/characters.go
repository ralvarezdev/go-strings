package count

// Numbers counts the number of numbers in a string
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
func Caps(s string) int {
	return Capital(s)
}

// Special counts the number of special characters in a string
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
func Alphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return count
}
