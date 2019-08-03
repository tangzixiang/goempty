package goempty

import (
	"reflect"
)

type Emp interface {
	Empty() bool
}

// IsZero 判断某个值是否零值
// 如果 value 实现了 `Emp` 则会返回 Empty 方法的返回值
func IsZero(value interface{}) bool {
	if value == nil {
		return true
	}

	if emptier, ok := value.(Emp); ok {
		return emptier.Empty()
	}

	switch v := value.(type) {
	case string:
		return IsStrEmpty(v)
	case int8:
		return IsInt8Empty(v)
	case int16:
		return IsInt16Empty(v)
	case int32:
		return IsInt32Empty(v)
	case int64:
		return IsInt64Empty(v)
	case uint8:
		return IsUint8Empty(v)
	case uint16:
		return IsUint16Empty(v)
	case uint32:
		return IsUint32Empty(v)
	case uint64:
		return IsUint64Empty(v)
	case float32:
		return IsFloat32Empty(v)
	case float64:
		return IsFloat64Empty(v)
	case bool:
		return v == false
	}
	
	kind := reflect.TypeOf(value).Kind()
	if kind == reflect.Map || kind == reflect.Func || kind == reflect.Chan || kind == reflect.Ptr  || kind == reflect.Slice{
		v := reflect.ValueOf(value)
		if v.IsNil()|| v.IsValid(){
			return true
		}

		return false
	}

	return reflect.Zero(reflect.TypeOf(value)).Interface() == value
}

// IsNotZero 判断某个值是否非零值
// 如果 value 实现了 `Emp` 则会返回 Empty 方法的返回值
func IsNotZero(value interface{}) bool {
	return !IsZero(value)
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
	return value, value == int8(0)
}

// IsEmptyInt16 是否数值为 0
func IsEmptyInt16(value int16) (int16, bool) {
	return value, value == int16(0)
}

// IsEmptyInt32 是否数值为 0
func IsEmptyInt32(value int32) (int32, bool) {
	return value, value == int32(0)
}

// IsEmptyInt64 是否数值为 0
func IsEmptyInt64(value int64) (int64, bool) {
	return value, value&int64(0) == 0
}

// IsEmptyFloat32 是否数值为 0
func IsEmptyFloat32(value float32) (float32, bool) {
	return value, value == float32(0)
}

// IsEmptyFloat64 是否数值为 0
func IsEmptyFloat64(value float64) (float64, bool) {
	return value, value == float64(0)
}

// IsEmptyUInt 是否数值为 0
func IsEmptyUInt(value uint) (uint, bool) {
	return value, value == uint(0)
}

// IsEmptyUInt8 是否数值为 0
func IsEmptyUInt8(value uint8) (uint8, bool) {
	return value, value == uint8(0)
}

// IsEmptyUInt16 是否数值为 0
func IsEmptyUInt16(value uint16) (uint16, bool) {
	return value, value == uint16(0)
}

// IsEmptyUInt32 是否数值为 0
func IsEmptyUInt32(value uint32) (uint32, bool) {
	return value, value == uint32(0)
}

// IsEmptyUInt64 是否数值为 0
func IsEmptyUInt64(value uint64) (uint64, bool) {
	return value, value == uint64(0)
}

// IsEmptyByte 是否数值为 0
func IsEmptyByte(value byte) (byte, bool) {
	return value, value == byte(0)
}

// IsStrEmpty 是否空字符串
func IsStrEmpty(str string) bool {
	return str == ""
}

// IsIntEmpty 是否数值为 0
func IsIntEmpty(value int) bool {
	return value == 0
}

// IsInt8Empty 是否数值为 0
func IsInt8Empty(value int8) bool {
	return value == int8(0)
}

// IsInt16Empty 是否数值为 0
func IsInt16Empty(value int16) bool {
	return value == int16(0)
}

// IsInt32Empty 是否数值为 0
func IsInt32Empty(value int32) bool {
	return value == int32(0)
}

// IsInt64Empty 是否数值为 0
func IsInt64Empty(value int64) bool {
	return value == int64(0)
}

// IsUintEmpty 是否数值为 0
func IsUintEmpty(value uint) bool {
	return value == uint(0)
}

// IsUint8Empty 是否数值为 0
func IsUint8Empty(value uint8) bool {
	return value == uint8(0)
}

// IsUint16Empty 是否数值为 0
func IsUint16Empty(value uint16) bool {
	return value == uint16(0)
}

// IsUint32Empty 是否数值为 0
func IsUint32Empty(value uint32) bool {
	return value == uint32(0)
}

// IsUint64Empty 是否数值为 0
func IsUint64Empty(value uint64) bool {
	return value == uint64(0)
}

// IsByteEmpty 是否数值为 0
func IsByteEmpty(value byte) bool {
	return value == byte(0)
}

// IsFloat32Empty 是否数值为 0
func IsFloat32Empty(value float32) bool {
	return value == float32(0)
}

// IsFloat64Empty 是否数值为 0
func IsFloat64Empty(value float64) bool {
	return value == float64(0)
}
