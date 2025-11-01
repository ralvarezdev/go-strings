package json

import (
	"fmt"
	"reflect"
	"strings"

	gostrings "github.com/ralvarezdev/go-strings"
)

// GetJSONTag returns the JSON tag for a given struct field
//
// Parameters:
//
// - structField: The struct field
// - fieldName: The name of the struct field
//
// Returns:
//
// - string: JSON tag
// - error: error if the field does not exist
func GetJSONTag(structField *reflect.StructField, fieldName string) (string, error) {
	// Check if the struct field is nil
	if structField == nil {
		return "", gostrings.ErrNilStructField
	}

	// Get the field from the struct field
	jsonTag := structField.Tag.Get(JSONTag)
	if jsonTag == "" {
		return "", fmt.Errorf(ErrJSONTagNotFound, fieldName)
	}
	return jsonTag, nil
}

// GetJSONTagName returns the JSON tag name for a given struct field name
//
// Parameters:
//
// - jsonTag: The JSON tag of the struct field
// - fieldName: The name of the struct field
//
// Returns:
//
// - string: JSON tag name
// - error: error if the JSON tag is empty
func GetJSONTagName(jsonTag, fieldName string) (string, error) {
	// Get the field name from the JSON tag
	tagParts := strings.Split(jsonTag, ",")
	if len(tagParts) == 0 {
		return "", fmt.Errorf(ErrEmptyJSONTag, fieldName)
	}
	return tagParts[0], nil
}
