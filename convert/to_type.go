package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt sets the destination to the value of the string as an int.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to an int where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to int.
func ToInt(value string, dest interface{}) error {
	destPtr, ok := dest.(*int)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Int.String())
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf(ErrParseIntFailed, value)
	}
	*destPtr = i
	return nil
}

// ToInt8 sets the destination to the value of the string as an int8.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to an int8 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to int8.
func ToInt8(value string, dest interface{}) error {
	destPtr, ok := dest.(*int8)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Int8.String())
	}

	i, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return fmt.Errorf(ErrParseInt8Failed, value)
	}
	*destPtr = int8(i)
	return nil
}

// ToInt16 sets the destination to the value of the string as an int16.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to an int16 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to int16.
func ToInt16(value string, dest interface{}) error {
	destPtr, ok := dest.(*int16)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Int16.String())
	}

	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return fmt.Errorf(ErrParseInt16Failed, value)
	}
	*destPtr = int16(i)
	return nil
}

// ToInt32 sets the destination to the value of the string as an int32.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to an int32 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to int32.
func ToInt32(value string, dest interface{}) error {
	destPtr, ok := dest.(*int32)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Int32.String())
	}

	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return fmt.Errorf(ErrParseInt32Failed, value)
	}
	*destPtr = int32(i)
	return nil
}

// ToInt64 sets the destination to the value of the string as an int64.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to an int64 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to int64.
func ToInt64(value string, dest interface{}) error {
	destPtr, ok := dest.(*int64)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Int64.String())
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fmt.Errorf(ErrParseInt64Failed, value)
	}
	*destPtr = i
	return nil
}

// ToUint sets the destination to the value of the string as a uint.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a uint where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to uint.
func ToUint(value string, dest interface{}) error {
	destPtr, ok := dest.(*uint)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Uint.String())
	}

	i, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return fmt.Errorf(ErrParseUintFailed, value)
	}
	*destPtr = uint(i)
	return nil
}

// ToUint8 sets the destination to the value of the string as a uint8.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a uint8 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to uint8.
func ToUint8(value string, dest interface{}) error {
	destPtr, ok := dest.(*uint8)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Uint8.String())
	}

	i, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return fmt.Errorf(ErrParseUint8Failed, value)
	}
	*destPtr = uint8(i)
	return nil
}

// ToUint16 sets the destination to the value of the string as a uint16.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a uint16 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to uint16.
func ToUint16(value string, dest interface{}) error {
	destPtr, ok := dest.(*uint16)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Uint16.String())
	}

	u, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return fmt.Errorf(ErrParseUint16Failed, value)
	}
	*destPtr = uint16(u)
	return nil
}

// ToUint32 sets the destination to the value of the string as a uint32.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a uint32 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to uint32.
func ToUint32(value string, dest interface{}) error {
	destPtr, ok := dest.(*uint32)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Uint32.String())
	}

	u, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return fmt.Errorf(ErrParseUint32Failed, value)
	}
	*destPtr = uint32(u)
	return nil
}

// ToUint64 sets the destination to the value of the string as a uint64.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a uint64 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to uint64.
func ToUint64(value string, dest interface{}) error {
	destPtr, ok := dest.(*uint64)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Uint64.String())
	}

	u, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fmt.Errorf(ErrParseUint64Failed, value)
	}
	*destPtr = u
	return nil
}

// ToFloat32 sets the destination to the value of the string as a float32.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a float32 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to float32.
func ToFloat32(value string, dest interface{}) error {
	destPtr, ok := dest.(*float32)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Float32.String())
	}

	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return fmt.Errorf(ErrParseFloat32Failed, value)
	}
	*destPtr = float32(f)
	return nil
}

// ToFloat64 sets the destination to the value of the string as a float64.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a float64 where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to float64.
func ToFloat64(value string, dest interface{}) error {
	destPtr, ok := dest.(*float64)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Float64.String())
	}

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf(ErrParseFloat64Failed, value)
	}
	*destPtr = f
	return nil
}

// ToBool sets the destination to the value of the string as a bool.
//
// Parameters:
//
//	value - the string to convert
//	dest - pointer to a bool where the result will be stored
//
// Returns:
//
//	An error if the conversion fails or the destination is not a pointer to bool.
func ToBool(value string, dest interface{}) error {
	destPtr, ok := dest.(*bool)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.Bool.String())
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf(ErrParseBoolFailed, value)
	}
	*destPtr = b
	return nil
}

// ToString sets the destination to the value of the string.
//
// Parameters:
//
//	value - the string to assign
//	dest - pointer to a string where the value will be stored
//
// Returns:
//
//	An error if the destination is not a pointer to string.
func ToString(value string, dest interface{}) error {
	destPtr, ok := dest.(*string)
	if !ok {
		return fmt.Errorf(ErrInvalidDestinationType, reflect.String.String())
	}

	*destPtr = value
	return nil
}
