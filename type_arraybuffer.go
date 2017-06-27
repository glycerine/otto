package otto

import (
	"encoding/binary"
	"math"
)

type _arrayBufferObject struct {
	buffer []byte
}

func (runtime *_runtime) newArrayBufferObject(length uint32) *_object {
	self := runtime.newObject()
	self.class = "ArrayBuffer"
	self.defineProperty("byteLength", toValue_uint32(length), 0100, false)
	self.objectClass = _classObject
	self.value = _arrayBufferObject{make([]byte, length)}
	return self
}

func (runtime *_runtime) newArrayBufferObjectOf(bytes []byte) *_object {
	self := runtime.newObject()
	self.class = "ArrayBuffer"
	self.defineProperty("byteLength", toValue_uint32(uint32(len(bytes))), 0100, false)
	self.objectClass = _classObject
	self.value = _arrayBufferObject{bytes}
	return self
}

func (self *_object) arrayBufferValue() _arrayBufferObject {
	value, _ := self.value.(_arrayBufferObject)
	return value
}

func arrayBufferObjectOf(rt *_runtime, _obj *_object) _arrayBufferObject {
	if _obj == nil || _obj.class != "ArrayBuffer" {
		panic(rt.panicTypeError())
	}
	return _obj.arrayBufferValue()
}

func getValueFromArrayBuffer(arrayBuffer _arrayBufferObject, byteIndex uint32, kind typedArrayKind) Value {
	byteLength := len(arrayBuffer.buffer)
	if byteIndex < 0 || byteIndex >= uint32(byteLength) {
		return UndefinedValue()
	}

	elementSize := typedArrayElementSize(kind)
	rawValue := arrayBuffer.buffer[byteIndex : byteIndex+elementSize]

	if kind == typedArrayKindFloat32 {
		intValue := binary.LittleEndian.Uint32(rawValue)
		value := math.Float32frombits(intValue)
		if math.IsNaN(float64(value)) {
			return NaNValue()
		}
		return toValue_float32(value)
	} else if kind == typedArrayKindFloat64 {
		intValue := binary.LittleEndian.Uint64(rawValue)
		value := math.Float64frombits(intValue)
		if math.IsNaN(value) {
			return NaNValue()
		}
		return toValue_float64(value)
	} else if kind == typedArrayKindInt8 {
		return toValue_int8(int8(rawValue[0]))
	} else if kind == typedArrayKindUint8 || kind == typedArrayKindUint8Clamped {
		return toValue_uint8(rawValue[0])
	} else if kind == typedArrayKindInt16 {
		value := binary.LittleEndian.Uint16(rawValue)
		return toValue_int16(int16(value))
	} else if kind == typedArrayKindUint16 {
		value := binary.LittleEndian.Uint16(rawValue)
		return toValue_uint16(value)
	} else if kind == typedArrayKindInt32 {
		value := binary.LittleEndian.Uint32(rawValue)
		return toValue_int32(int32(value))
	} else if kind == typedArrayKindUint32 {
		value := binary.LittleEndian.Uint32(rawValue)
		return toValue_uint32(value)
	}

	return UndefinedValue()
}

func setValueInArrayBuffer(arrayBuffer _arrayBufferObject, byteIndex uint32, kind typedArrayKind, value Value) {
	elementSize := typedArrayElementSize(kind)
	slice := arrayBuffer.buffer[byteIndex : byteIndex+elementSize]

	if kind == typedArrayKindFloat32 {
		bits := math.Float32bits(float32(value.number().float64))
		binary.LittleEndian.PutUint32(slice, bits)
	} else if kind == typedArrayKindFloat64 {
		bits := math.Float64bits(value.number().float64)
		binary.LittleEndian.PutUint64(slice, bits)
	} else if kind == typedArrayKindInt8 || kind == typedArrayKindUint8 {
		slice[0] = byte(value.number().int64 % 256)
	} else if kind == typedArrayKindUint8Clamped {
		intValue := value.number().int64
		if intValue > 255 {
			intValue = 255
		} else if intValue < 0 {
			intValue = 0
		}
		slice[0] = uint8(intValue)
	} else if kind == typedArrayKindInt16 || kind == typedArrayKindUint16 {
		n := value.number().int64
		if kind == typedArrayKindUint16 && n < 0 {
			maxValue := int64(math.Pow(2, float64(elementSize*8)))
			n = maxValue + n
		}
		binary.LittleEndian.PutUint16(slice, uint16(n))
	} else if kind == typedArrayKindInt32 || kind == typedArrayKindUint32 {
		n := value.number().int64
		if kind == typedArrayKindUint32 && n < 0 {
			maxValue := int64(math.Pow(2, float64(elementSize*8)))
			n = maxValue + n
		}
		binary.LittleEndian.PutUint32(slice, uint32(n))
	}
}
