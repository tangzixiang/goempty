package goempty

import (
	"reflect"
)

// IsZero 判断某个值是否零值
func IsZero(value interface{}) bool {
	if value == nil {
		return false
	}
	return reflect.Zero(reflect.TypeOf(value)).Interface() == value
}

// IsEmptyStr 是否空字符串
func IsEmptyStr(str string) (string, bool) {
	return str, str == ""
}

// IsEmptyInt 是否数值为 0
func IsEmptyInt(value int) (int, bool) {
	return value, value == 0
}

// IsEmptyInt8 是否数值为 0
func IsEmptyInt8(value int8) (int8, bool) {
	return value, value == 0
}
// IsEmptyInt16 是否数值为 0
func IsEmptyInt16(value int16) (int16, bool) {
	return value, value == 0
}

// IsEmptyInt32 是否数值为 0
func IsEmptyInt32(value int32) (int32, bool) {
	return value, value == 0
}

// IsEmptyInt64 是否数值为 0
func IsEmptyInt64(value int64) (int64, bool) {
	return value, value == 0
}

// IsEmptyUInt 是否数值为 0
func IsEmptyUInt(value uint) (uint, bool) {
	return value, value == 0
}

// IsEmptyUInt8 是否数值为 0
func IsEmptyUInt8(value uint8) (uint8, bool) {
	return value, value == 0
}

// IsEmptyUInt16 是否数值为 0
func IsEmptyUInt16(value uint16) (uint16, bool) {
	return value, value == 0
}

// IsEmptyUInt32 是否数值为 0
func IsEmptyUInt32(value uint32) (uint32, bool) {
	return value, value == 0
}

// IsEmptyUInt64 是否数值为 0
func IsEmptyUInt64(value uint64) (uint64, bool) {
	return value, value == 0
}

// IsEmptyByte 是否数值为 0
func IsEmptyByte(value byte) (byte, bool) {
	return value, value == 0
}