package goempty

type Def interface {
	Default() interface{}
}

// Default 判断是否内容为零值，是则返回默认值
// 如果 old 实现了 `Def` 则会返回 Default 方法的返回值
func Default(old interface{}, defaultV ...interface{}) interface{} {

	if IsZero(old) && len(defaultV) >0 {
		return defaultV[0]
	}

	if d, ok := old.(Def); ok {
		return d.Default()
	}

	return nil
}

// DefaultStr 判断是否内容为零值，是则返回默认值
func DefaultStr(old, defaultV string) string {
	if IsStrEmpty(old) {
		return defaultV
	}

	return old
}

// DefaultInt 判断是否内容为零值，是则返回默认值
func DefaultInt(old, defaultV int) int {
	if IsIntEmpty(old) {
		return defaultV
	}

	return old
}

// DefaultInt8 判断是否内容为零值，是则返回默认值
func DefaultInt8(old, defaultV int8) int8 {
	if IsInt8Empty(old) {
		return defaultV
	}

	return old
}

// DefaultInt16 判断是否内容为零值，是则返回默认值
func DefaultInt16(old, defaultV int16) int16 {
	if IsInt16Empty(old) {
		return defaultV
	}

	return old
}

// DefaultInt32 判断是否内容为零值，是则返回默认值
func DefaultInt32(old, defaultV int32) int32 {
	if IsInt32Empty(old) {
		return defaultV
	}

	return old
}

// DefaultInt64 判断是否内容为零值，是则返回默认值
func DefaultInt64(old, defaultV int64) int64 {
	if IsInt64Empty(old) {
		return defaultV
	}

	return old
}

// DefaultUint 判断是否内容为零值，是则返回默认值
func DefaultUint(old, defaultV uint) uint {
	if IsUintEmpty(old) {
		return defaultV
	}

	return old
}

// DefaultUint8 判断是否内容为零值，是则返回默认值
func DefaultUint8(old, defaultV uint8) uint8 {
	if IsUint8Empty(old) {
		return defaultV
	}

	return old
}

// DefaultUint16 判断是否内容为零值，是则返回默认值
func DefaultUint16(old, defaultV uint16) uint16 {
	if IsUint16Empty(old) {
		return defaultV
	}

	return old
}

// DefaultInt32 判断是否内容为零值，是则返回默认值
func DefaultUint32(old, defaultV uint32) uint32 {
	if IsUint32Empty(old) {
		return defaultV
	}

	return old
}

// DefaultUint64 判断是否内容为零值，是则返回默认值
func DefaultUint64(old, defaultV uint64) uint64 {
	if IsUint64Empty(old) {
		return defaultV
	}

	return old
}

// DefaultFloat32 判断是否内容为零值，是则返回默认值
func DefaultFloat32(old, defaultV float32) float32 {
	if IsFloat32Empty(old) {
		return defaultV
	}

	return old
}

// DefaultFloat64 判断是否内容为零值，是则返回默认值
func DefaultFloat64(old, defaultV float64) float64 {
	if IsFloat64Empty(old) {
		return defaultV
	}

	return old
}

// DefaultByte 判断是否内容为零值，是则返回默认值
func DefaultByte(old, defaultV byte) byte {
	if IsByteEmpty(old) {
		return defaultV
	}

	return old
}
