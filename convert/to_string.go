package convert

import (
	"fmt"
	"strings"
	"time"
)

var (
	validUnits = []string{"ns", "us", "ms", "s", "m", "h", "d"}
)

// ErrorArrayToStringArray maps an array of errors to an array of strings
func ErrorArrayToStringArray(errorArray *[]error) *[]string {
	if errorArray == nil || len(*errorArray) == 0 {
		return nil
	}

	// Map the errors to strings
	stringArray := make([]string, len(*errorArray))
	for i, err := range *errorArray {
		stringArray[i] = err.Error()
	}
	return &stringArray
}

// PrettyDuration formats a duration without zero units
func PrettyDuration(d time.Duration, minUnit string) (string, error) {
	if d == 0 {
		return "0s", nil
	}

	// Check if the duration is negative
	if d < 0 {
		d = -d
	}

	// Check if the minimum unit is valid
	isValid := false
	for _, unit := range validUnits {
		if minUnit == unit {
			isValid = true
			break
		}
	}
	if !isValid {
		return "", fmt.Errorf(
			ErrInvalidDurationUnit,
			strings.Join(validUnits, ", "),
		)
	}

	var result []string
	includeNanoseconds := minUnit == "ns"
	includeMicroseconds := minUnit == "us" || includeNanoseconds
	includeMilliseconds := minUnit == "ms" || includeMicroseconds
	includeSeconds := minUnit == "s" || includeMilliseconds
	includeMinutes := minUnit == "m" || includeSeconds
	includeHours := minUnit == "h" || includeMinutes
	includeDays := minUnit == "d" || includeHours

	years := int64(d.Hours()) / 24 / 365
	if years > 0 {
		result = append(result, fmt.Sprintf("%dy", years))
	}

	if includeDays {
		days := int64(d.Hours()) / 24 % 365
		if days > 0 {
			result = append(result, fmt.Sprintf("%dd", days))
		}
	}

	if includeHours {
		hours := int64(d.Hours()) % 24
		if hours > 0 {
			result = append(result, fmt.Sprintf("%dh", hours))
		}
	}

	if includeMinutes {
		minutes := int64(d.Minutes()) % 60
		if minutes > 0 {
			result = append(result, fmt.Sprintf("%dm", minutes))
		}
	}

	if includeSeconds {
		seconds := int64(d.Seconds()) % 60
		if seconds > 0 {
			result = append(result, fmt.Sprintf("%ds", seconds))
		}
	}

	if includeMilliseconds {
		milliseconds := d.Milliseconds() % 1000
		if milliseconds > 0 {
			result = append(result, fmt.Sprintf("%dms", milliseconds))
		}
	}

	if includeMicroseconds {
		microseconds := d.Microseconds() % 1000
		if microseconds > 0 {
			result = append(result, fmt.Sprintf("%dus", microseconds))
		}
	}

	if includeNanoseconds {
		nanoseconds := d.Nanoseconds() % 1000
		if nanoseconds > 0 {
			result = append(result, fmt.Sprintf("%dns", nanoseconds))
		}
	}

	return strings.Join(result, " "), nil
}
