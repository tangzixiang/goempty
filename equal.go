package goempty

import (
	"reflect"
)

type Eq interface {
	Equal(interface{}) bool
}

// Equal 判断两者是否一致
// 如果 self 实现了 `Eq` 则会调用 self.Eq(comp)
// 如果 comp 实现了 `Eq` 则会调用 comp.Eq(self)
// 如果两者均未实现 `Eq` 则会调用 `reflect.DeepEqual` 进行进一步的判断
func Equal(self interface{}, comp interface{}) bool {

	if self == nil || comp == nil {
		return false
	}

	if e, ok := self.(Eq); ok {
		return e.Equal(comp)
	}

	if e, ok := comp.(Eq); ok {
		return e.Equal(self)
	}

	return reflect.DeepEqual(self, comp)
}

// EqualStr 是否空字符串
func EqualStr(str1, str2 string) bool {
	return str1 == str2
}

// EqualInt 两者比较是否相等
func EqualInt(value1, value2 int) bool {
	return value1 == value2
}

// EqualInt8 两者比较是否相等
func EqualInt8(value1, value2 int8) bool {
	return value1 == value2
}

// EqualInt16 两者比较是否相等
func EqualInt16(value1, value2 int16) bool {
	return value1 == value2
}

// EqualInt32 两者比较是否相等
func EqualInt32(value1, value2 int32) bool {
	return value1 == value2
}

// EqualInt64 两者比较是否相等
func EqualInt64(value1, value2 int64) bool {
	return value1 == value2
}

// EqualUint 两者比较是否相等
func EqualUint(value1, value2 uint) bool {
	return value1 == value2
}

// EqualUint8 两者比较是否相等
func EqualUint8(value1, value2 uint8) bool {
	return value1 == value2
}

// EqualUint16 两者比较是否相等
func EqualUint16(value1, value2 uint16) bool {
	return value1 == value2
}

// EqualUint32 两者比较是否相等
func EqualUint32(value1, value2 uint32) bool {
	return value1 == value2
}

// EqualUint64 两者比较是否相等
func EqualUint64(value1, value2 uint64) bool {
	return value1 == value2
}

// EqualByte 两者比较是否相等
func EqualByte(value1, value2 byte) bool {
	return value1 == value2
}

// EqualFloat32 两者比较是否相等
func EqualFloat32(value1, value2 float32) bool {
	return value1 == value2
}

// EqualFloat64 两者比较是否相等
func EqualFloat64(value1, value2 float64) bool {
	return value1 == value2
}
