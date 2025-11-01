package protobuf

import (
	"fmt"
	"reflect"
	"strings"

	gostrings "github.com/ralvarezdev/go-strings"
)

// GetProtobufTag gets the Protobuf tag number for a given struct field name
//
// Parameters:
//
// - structField: The struct field
// - fieldName: The name of the struct field
//
// Returns:
//
// - string: Protobuf tag
// - error: error if the Protobuf tag is not found
func GetProtobufTag(structField *reflect.StructField, fieldName string) (string, error) {
	protobufTag := structField.Tag.Get(ProtobufTag)
	if protobufTag == "" {
		return "", fmt.Errorf(ErrProtobufTagNotFound, fieldName)
	}
	return protobufTag, nil
}

// IsProtobufOneOfField checks if the struct field is a Protobuf oneof field
//
// Parameters:
//
// - structField: The struct field
//
// Returns:
//
// - bool: true if the struct field is a Protobuf oneof field
// - error: error if the struct field is nil
func IsProtobufOneOfField(structField *reflect.StructField) (bool, error) {
	// Check if the struct type is nil
	if structField == nil {
		return false, gostrings.ErrNilStructField
	}
	return structField.Tag.Get(ProtobufOneOfTag) != "", nil
}

// GetProtobufTagName returns the Protobuf tag name for a given struct field name
//
// Parameters:
//
//   - protobufTag: The Protobuf tag of the struct field
//   - fieldName: The name of the struct field
//
// Returns:
//
//   - string: Protobuf tag name
func GetProtobufTagName(protobufTag, fieldName string) (string, error) {
	for _, part := range strings.Split(protobufTag, ",") {
		if strings.HasPrefix(part, ProtobufNamePrefix) {
			return strings.TrimPrefix(part, ProtobufNamePrefix), nil
		}
	}
	return "", fmt.Errorf(ErrProtobufTagNameNotFound, fieldName)
}

// IsProtobufFieldOptional returns true if the Protobuf field is optional
//
// Parameters:
//
//   - protobufTag: The Protobuf tag of the struct field
//
// Returns:
//
//   - bool: true if the Protobuf field is optional
func IsProtobufFieldOptional(protobufTag string) bool {
	return strings.Contains(
		protobufTag,
		ProtobufOneOf,
	)
}

// IsProtobufGeneratedField checks if the field is a Protobuf generated field
//
// Parameters:
//
//   - fieldName: The name of the struct field
//
// Returns:
//
// - bool: true if the field is a Protobuf generated field
func IsProtobufGeneratedField(fieldName string) bool {
	return fieldName == State || fieldName == SizeCache || fieldName == UnknownFields
}
