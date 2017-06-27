package otto

import "encoding/binary"

type typedArrayKind byte

const (
	typedArrayKindInt8 typedArrayKind = iota
	typedArrayKindInt16
	typedArrayKindInt32
	typedArrayKindUint8
	typedArrayKindUint8Clamped
	typedArrayKindUint16
	typedArrayKindUint32
	typedArrayKindFloat32
	typedArrayKindFloat64
)

func typedArrayElementSize(kind typedArrayKind) uint32 {
	switch kind {
	case typedArrayKindInt8:
		return 1
	case typedArrayKindInt16:
		return 2
	case typedArrayKindInt32:
		return 4
	case typedArrayKindUint8:
		return 1
	case typedArrayKindUint8Clamped:
		return 1
	case typedArrayKindUint16:
		return 2
	case typedArrayKindUint32:
		return 4
	case typedArrayKindFloat32:
		return 4
	case typedArrayKindFloat64:
		return 8
	}

	panic("Unknown typedArrayKind")
}

func typedArrayKindToClass(kind typedArrayKind) string {
	switch kind {
	case typedArrayKindInt8:
		return "Int8Array"
	case typedArrayKindInt16:
		return "Int16Array"
	case typedArrayKindInt32:
		return "Int32Array"
	case typedArrayKindUint8:
		return "Uint8Array"
	case typedArrayKindUint8Clamped:
		return "Uint8ClampedArray"
	case typedArrayKindUint16:
		return "Uint16Array"
	case typedArrayKindUint32:
		return "Uint32Array"
	case typedArrayKindFloat32:
		return "Float32Array"
	case typedArrayKindFloat64:
		return "Float64Array"
	}

	panic("Unknown typedArrayKind")
}

func typedArrayKindFromClass(className string) typedArrayKind {
	switch className {
	case "Int8Array":
		return typedArrayKindInt8
	case "Int16Array":
		return typedArrayKindInt16
	case "Int32Array":
		return typedArrayKindInt32
	case "Uint8Array":
		return typedArrayKindUint8
	case "Uint8ClampedArray":
		return typedArrayKindUint8Clamped
	case "Uint16Array":
		return typedArrayKindUint16
	case "Uint32Array":
		return typedArrayKindUint32
	case "Float32Array":
		return typedArrayKindFloat32
	case "Float64Array":
		return typedArrayKindFloat64
	}

	panic("Unknown typedArrayKind")
}

func typedArrayElementSizeFromClass(className string) uint32 {
	switch className {
	case "Int8Array":
		return 1
	case "Int16Array":
		return 2
	case "Int32Array":
		return 4
	case "Uint8Array":
		return 1
	case "Uint8ClampedArray":
		return 1
	case "Uint16Array":
		return 2
	case "Uint32Array":
		return 4
	case "Float32Array":
		return 4
	case "Float64Array":
		return 8
	}

	panic("Unknown typedArrayKind")
}

func getInt8FromBytes(b []byte) int8 {
	return int8(b[0])
}

func getUint8FromBytes(b []byte) uint8 {
	return uint8(b[0])
}

func getInt16FromBytes(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(b))
}

func getUint16FromBytes(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func getInt32FromBytes(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}

func getUint32FromBytes(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func (runtime *_runtime) newTypedArrayObject(length uint32, buffer *_object, byteOffset uint32, byteLength uint32, kind typedArrayKind) *_object {
	self := runtime.newObject()
	self.class = typedArrayKindToClass(kind)
	self.objectClass = _classTypedArray
	self.defineProperty("length", toValue_uint32(length), 0100, false)
	self.defineProperty("buffer", toValue_object(buffer), 0000, false)
	self.defineProperty("byteOffset", toValue_uint32(byteOffset), 0000, false)
	self.defineProperty("byteLength", toValue_uint32(byteLength), 0000, false)
	return self
}

func isTypedArray(object *_object) bool {
	if object == nil {
		return false
	}

	switch object.class {
	case "Int8Array",
		"Int16Array",
		"Int32Array",
		"Uint8Array",
		"Uint8ClampedArray",
		"Uint16Array",
		"Uint32Array",
		"Float32Array",
		"Float64Array":
		return true
	default:
		return false
	}
}

func typedArrayGetLengthAndByteLength(self *_object) (uint32, uint32) {
	lengthProperty := self.getOwnProperty("length")
	lengthValue, valid := lengthProperty.value.(Value)
	if !valid {
		panic("TypedArray.length != Value{}")
	}

	byteLengthProperty := self.getOwnProperty("byteLength")
	byteLengthValue, valid := byteLengthProperty.value.(Value)
	if !valid {
		panic("TypedArray.byteLength != Value{}")
	}

	length := lengthValue.value.(uint32)
	byteLength := byteLengthValue.value.(uint32)

	return length, byteLength
}

func _typedArrayPut(self *_object, index int64, value Value) {
	length, byteLength := typedArrayGetLengthAndByteLength(self)
	elementSize := byteLength / length
	arrayBuffer := self.get("buffer")._object().arrayBufferValue()
	byteOffset := uint32(self.get("byteOffset").number().int64)
	byteIndex := byteOffset + uint32(index)*elementSize
	setValueInArrayBuffer(arrayBuffer, byteIndex, typedArrayKindFromClass(self.class), value)
}

func typedArrayPut(self *_object, name string, value Value, throw bool) {
	if index := stringToArrayIndex(name); index >= 0 {
		_typedArrayPut(self, index, value)
		return
	}

	objectPut(self, name, value, throw)
	return
}

func _typedArrayGet(self *_object, index int64) Value {
	length, byteLength := typedArrayGetLengthAndByteLength(self)
	elementSize := byteLength / length
	arrayBuffer := self.get("buffer")._object().arrayBufferValue()
	byteOffset := uint32(self.get("byteOffset").number().int64)
	byteIndex := byteOffset + uint32(index)*elementSize
	return getValueFromArrayBuffer(arrayBuffer, byteIndex, typedArrayKindFromClass(self.class))
}

func typedArrayGet(self *_object, name string) Value {
	if index := stringToArrayIndex(name); index >= 0 {
		return _typedArrayGet(self, index)
	}

	return objectGet(self, name)
}

func typedArrayGetOwnProperty(self *_object, name string) *_property {
	if index := stringToArrayIndex(name); index >= 0 {
		value := _typedArrayGet(self, index)
		return &_property{value, 0000}
	}

	return objectGetOwnProperty(self, name)
}
