package convert

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
