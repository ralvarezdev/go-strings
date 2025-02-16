package convert

var (
	ErrInvalidDestinationType = "invalid destination type, expected: '%s'"
	ErrParseIntFailed         = "cannot parse '%s', expected integer"
	ErrParseInt8Failed        = "cannot parse '%s', expected integer of 8 bits"
	ErrParseInt16Failed       = "cannot parse '%s', expected integer of 16 bits"
	ErrParseInt32Failed       = "cannot parse '%s', expected integer of 32 bits"
	ErrParseInt64Failed       = "cannot parse '%s', expected integer of 64 bits"
	ErrParseUintFailed        = "cannot parse '%s', expected unsigned integer"
	ErrParseUint8Failed       = "cannot parse '%s', expected unsigned integer of 8 bits"
	ErrParseUint16Failed      = "cannot parse '%s', expected unsigned integer of 16 bits"
	ErrParseUint32Failed      = "cannot parse '%s', expected unsigned integer of 32 bits"
	ErrParseUint64Failed      = "cannot parse '%s', expected unsigned integer of 64 bits"
	ErrParseFloat32Failed     = "cannot parse '%s', expected float of 32 bits"
	ErrParseFloat64Failed     = "cannot parse '%s', expected float of 64 bits"
	ErrParseBoolFailed        = "cannot parse '%s', expected boolean"
	ErrInvalidDurationUnit    = "invalid duration unit, expected: '%s'"
)
