package convert

import (
	"strconv"
	"strings"
)

type (
	ToType func(string) (interface{}, error)
)

// ToInt converts a string to an int
func ToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// ToInt8 converts a string to an int8
func ToInt8(value string) (int8, error) {
	i, err := strconv.ParseInt(value, 10, 8)
	return int8(i), err
}

// ToInt16 converts a string to an int16
func ToInt16(value string) (int16, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	return int16(i), err
}

// ToInt32 converts a string to an int32
func ToInt32(value string) (int32, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	return int32(i), err
}

// ToInt64 converts a string to an int64
func ToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

// ToUint converts a string to an uint
func ToUint(value string) (uint, error) {
	i, err := strconv.ParseUint(value, 10, 0)
	return uint(i), err
}

// ToUint8 converts a string to an uint8
func ToUint8(value string) (uint8, error) {
	i, err := strconv.ParseUint(value, 10, 8)
	return uint8(i), err
}

// ToUint16 converts a string to an uint16
func ToUint16(value string) (uint16, error) {
	i, err := strconv.ParseUint(value, 10, 16)
	return uint16(i), err
}

// ToUint32 converts a string to an uint32
func ToUint32(value string) (uint32, error) {
	i, err := strconv.ParseUint(value, 10, 32)
	return uint32(i), err
}

// ToUint64 converts a string to an uint64
func ToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

// ToFloat32 converts a string to a float32
func ToFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	return float32(f), err
}

// ToFloat64 converts a string to a float64
func ToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// ToBool converts a string to a bool
func ToBool(value string) (bool, error) {
	if value == strings.ToLower(strings.TrimSpace("true")) {
		return true, nil
	}
	return false, nil
}
