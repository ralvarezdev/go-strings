package json

import (
	"fmt"
	"reflect"
	"strings"
)

// GetJSONTag returns the JSON tag for a given struct field
//
// Parameters:
//
// - structFieldType: The struct type
// - fieldName: The name of the struct field
//
// Returns:
//
// - string: JSON tag
// - error: error if the field does not exist
func GetJSONTag(structType reflect.StructField, fieldName string) (string, error) {
	// Get the field from the struct type
	jsonTag := structType.Tag.Get(JSONTag)
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