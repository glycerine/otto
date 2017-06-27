package otto

func typedArray(kind typedArrayKind, runtime *_runtime, argumentList []Value) Value {
	var length uint32
	var byteLength uint32
	var byteOffset uint32
	var buffer *_object

	isFirstArgBuffer := false
	elementSize := typedArrayElementSize(kind)

	if len(argumentList) == 0 { // new TypedArray(); // new in ES2017
		buffer = runtime.newArrayBuffer(0)
	} else {
		firstArgument := argumentList[0]
		if firstArgument.IsNumber() { // new TypedArray(length);
			length = uint32(firstArgument.number().int64)
			byteLength = length * elementSize
			buffer = runtime.newArrayBuffer(byteLength)
		} else if firstArgument.IsTypedArray() { // new TypedArray(typedArray);
			obj := firstArgument._object()
			arrayBuffer := obj.get("buffer")._object().arrayBufferValue()
			byteLength := uint32(obj.get("byteLength").number().int64)
			length = byteLength / elementSize
			buffer = runtime.newArrayBuffer(byteLength)

			copy(buffer.arrayBufferValue().buffer, arrayBuffer.buffer)
		} else if firstArgument.isArrayBuffer() {
			buffer = firstArgument._object()
			length = uint32(buffer.get("length").number().int64)
			byteLength = uint32(buffer.get("byteLength").number().int64)
			isFirstArgBuffer = true
		} else if firstArgument.isArray() { // new TypedArray(object);
			arr := firstArgument._object()
			length = uint32(arr.get("length").number().int64)
			byteLength = length * elementSize
			buffer = runtime.newArrayBuffer(byteLength)

			var i int64
			for i = 0; i < int64(length); i++ {
				key := arrayIndexToString(i)
				value := arr.get(key)
				byteIndex := uint32(i) * elementSize
				setValueInArrayBuffer(buffer.arrayBufferValue(), byteIndex, kind, value)
			}
		}

		// new TypedArray(buffer, byteOffset, length);
		if len(argumentList) >= 2 && isFirstArgBuffer {
			secondArgument := argumentList[1]
			if secondArgument.IsNumber() {
				byteOffset = uint32(secondArgument.number().int64)
				if byteOffset < 0 || byteOffset%byteLength != 0 {
					runtime.panicRangeError("Invalid typed array length")
				}
			}
		}

		if isFirstArgBuffer {
			isThirdArgumentDefined := false
			bufferByteLength := uint32(buffer.get("byteLength").number().int64)

			if len(argumentList) >= 3 {
				thirdArgument := argumentList[2]
				if thirdArgument.IsNumber() {
					isThirdArgumentDefined = true
					length = uint32(thirdArgument.number().int64)
				}
			}

			if isThirdArgumentDefined {
				byteLength = length * elementSize
				if byteOffset+byteLength > bufferByteLength {
					runtime.panicRangeError("Invalid typed array length")
				}
			} else {
				if bufferByteLength%elementSize != 0 {
					runtime.panicRangeError("Invalid typed array length")
				}

				byteLength = bufferByteLength - byteOffset
				if byteLength < 0 {
					runtime.panicRangeError("Invalid typed array length")
				}
			}
		}
	}

	typedArray := runtime.newTypedArray(length, buffer, byteOffset, byteLength, kind)
	return toValue_object(typedArray)
}

func builtinInt8Array(call FunctionCall) Value {
	return typedArray(typedArrayKindInt8, call.runtime, call.ArgumentList)
}

func builtinNewInt8Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindInt8, self.runtime, argumentList)
}

func builtinInt16Array(call FunctionCall) Value {
	return typedArray(typedArrayKindInt16, call.runtime, call.ArgumentList)
}

func builtinNewInt16Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindInt16, self.runtime, argumentList)
}

func builtinInt32Array(call FunctionCall) Value {
	return typedArray(typedArrayKindInt32, call.runtime, call.ArgumentList)
}

func builtinNewInt32Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindInt32, self.runtime, argumentList)
}

func builtinUint8Array(call FunctionCall) Value {
	return typedArray(typedArrayKindUint8, call.runtime, call.ArgumentList)
}

func builtinNewUint8Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindUint8, self.runtime, argumentList)
}

func builtinUint8ClampedArray(call FunctionCall) Value {
	return typedArray(typedArrayKindUint8Clamped, call.runtime, call.ArgumentList)
}

func builtinNewUint8ClampedArray(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindUint8Clamped, self.runtime, argumentList)
}

func builtinUint16Array(call FunctionCall) Value {
	return typedArray(typedArrayKindUint16, call.runtime, call.ArgumentList)
}

func builtinNewUint16Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindUint16, self.runtime, argumentList)
}

func builtinUint32Array(call FunctionCall) Value {
	return typedArray(typedArrayKindUint32, call.runtime, call.ArgumentList)
}

func builtinNewUint32Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindUint32, self.runtime, argumentList)
}

func builtinFloat32Array(call FunctionCall) Value {
	return typedArray(typedArrayKindFloat32, call.runtime, call.ArgumentList)
}

func builtinNewFloat32Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindFloat32, self.runtime, argumentList)
}

func builtinFloat64Array(call FunctionCall) Value {
	return typedArray(typedArrayKindFloat64, call.runtime, call.ArgumentList)
}

func builtinNewFloat64Array(self *_object, argumentList []Value) Value {
	return typedArray(typedArrayKindFloat64, self.runtime, argumentList)
}

func builtinTypedArray_copyWithin(call FunctionCall) Value {
	return builtinArray_copyWithin(call)
}

func builtinTypedArray_entries(call FunctionCall) Value {
	panic("Unimplemented")
}

func builtinTypedArray_every(call FunctionCall) Value {
	return builtinArray_every(call)
}

func builtinTypedArray_fill(call FunctionCall) Value {
	return builtinArray_fill(call)
}

func builtinTypedArray_filter(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if callback := call.Argument(0); callback.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		byteLength := toUint32(thisObject.get("byteLength"))
		callThis := call.Argument(1)
		values := make([]Value, 0)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				value := thisObject.get(key)
				if callback.call(call.runtime, callThis, value, index, this).bool() {
					values = append(values, value)
				}
			}
		}

		kind := typedArrayKindFromClass(thisObject.class)
		buffer := call.runtime.newArrayBuffer(byteLength)
		typedArray := call.runtime.newTypedArray(uint32(len(values)), buffer, 0, byteLength, kind)
		for i, value := range values {
			k := arrayIndexToString(int64(i))
			typedArray.put(k, value, true)
		}
		return toValue_object(typedArray)
	}
	panic(call.runtime.panicTypeError())
}

func builtinTypedArray_find(call FunctionCall) Value {
	return builtinArray_find(call)
}

func builtinTypedArray_findIndex(call FunctionCall) Value {
	return builtinArray_findIndex(call)
}

func builtinTypedArray_forEach(call FunctionCall) Value {
	return builtinArray_forEach(call)
}

func builtinTypedArray_includes(call FunctionCall) Value {
	return builtinArray_includes(call)
}

func builtinTypedArray_indexOf(call FunctionCall) Value {
	return builtinArray_indexOf(call)
}

func builtinTypedArray_join(call FunctionCall) Value {
	return builtinArray_join(call)
}

func builtinTypedArray_keys(call FunctionCall) Value {
	panic("Unimplemented")
}

func builtinTypedArray_lastIndexOf(call FunctionCall) Value {
	return builtinArray_lastIndexOf(call)
}

func builtinTypedArray_map(call FunctionCall) Value {
	thisObject := call.thisObject()
	this := toValue_object(thisObject)
	if callback := call.Argument(0); callback.isCallable() {
		length := int64(toUint32(thisObject.get("length")))
		byteLength := toUint32(thisObject.get("byteLength"))
		callThis := call.Argument(1)
		values := make([]Value, length)
		for index := int64(0); index < length; index++ {
			if key := arrayIndexToString(index); thisObject.hasProperty(key) {
				values[index] = callback.call(call.runtime, callThis, thisObject.get(key), index, this)
			} else {
				values[index] = Value{}
			}
		}

		kind := typedArrayKindFromClass(thisObject.class)
		buffer := call.runtime.newArrayBuffer(byteLength)
		typedArray := call.runtime.newTypedArray(uint32(len(values)), buffer, 0, byteLength, kind)
		for i, value := range values {
			k := arrayIndexToString(int64(i))
			typedArray.put(k, value, true)
		}
		return toValue_object(typedArray)
	}
	panic(call.runtime.panicTypeError())
}

func builtinTypedArray_reduce(call FunctionCall) Value {
	return builtinArray_reduce(call)
}

func builtinTypedArray_reduceRight(call FunctionCall) Value {
	return builtinArray_reduceRight(call)
}

func builtinTypedArray_reverse(call FunctionCall) Value {
	return builtinArray_reverse(call)
}

func builtinTypedArray_set(call FunctionCall) Value {
	this := call.thisObject()
	if !isTypedArray(this) {
		panic(call.runtime.panicTypeError("this is not a typed array"))
	}

	array := call.Argument(0)
	src := array._object()
	if !isArray(src) && !isTypedArray(src) {
		panic(call.runtime.panicTypeError("invalid argument"))
	}

	offset := call.Argument(1)
	var targetOffset int64
	if !offset.IsUndefined() {
		targetOffset = offset.number().int64
	}

	if targetOffset < 0 {
		panic(call.runtime.panicRangeError("Start offset is negative"))
	}

	targetKind := typedArrayKindFromClass(this.class)
	targetBufferValue := this.get("buffer")
	targetBuffer := targetBufferValue._object().arrayBufferValue()
	targetLength := this.get("length").number().int64
	targetByteOffset := uint32(this.get("byteOffset").number().int64)
	targetElementSize := typedArrayElementSize(targetKind)

	srcLength := src.get("length").number().int64
	if srcLength+targetOffset > targetLength {
		panic(call.runtime.panicRangeError("invalid argument"))
	}

	targetByteIndex := uint32(targetOffset)*targetElementSize + targetByteOffset
	limit := targetByteIndex + targetElementSize*uint32(srcLength)

	if isArray(src) {
		var k int64
		for targetByteIndex < limit {
			pk := arrayIndexToString(k)
			value := src.get(pk)
			setValueInArrayBuffer(targetBuffer, targetByteIndex, targetKind, value)
			k += 1
			targetByteIndex += targetElementSize
		}
	} else if isTypedArray(src) {
		srcBufferValue := src.get("buffer")
		srcBuffer := srcBufferValue._object().arrayBufferValue()
		srcByteOffset := uint32(src.get("byteOffset").number().int64)
		srcKind := typedArrayKindFromClass(src.class)
		srcElementSize := typedArrayElementSize(srcKind)

		var srcByteIndex uint32
		if sameValue(srcBufferValue, targetBufferValue) {
			copy(srcBuffer.buffer, targetBuffer.buffer[srcByteOffset:])
		} else {
			srcByteIndex = srcByteOffset
		}

		if srcKind != targetKind {
			for targetByteIndex < limit {
				value := getValueFromArrayBuffer(srcBuffer, srcByteIndex, srcKind)
				setValueInArrayBuffer(targetBuffer, targetByteIndex, targetKind, value)
				srcByteIndex += srcElementSize
				targetByteIndex += targetElementSize
			}
		} else {
			// NOTE: If srcKind and targetKind are the same
			// the transfer must be performed in a manner
			// that preserves the bit-level encoding of the source data.
			for targetByteIndex < limit {
				value := getValueFromArrayBuffer(srcBuffer, srcByteIndex, typedArrayKindUint8)
				setValueInArrayBuffer(targetBuffer, targetByteIndex, typedArrayKindUint8, value)
				srcByteIndex += 1
				targetByteIndex += 1
			}
		}
	}

	return UndefinedValue()
}

func builtinTypedArray_slice(call FunctionCall) Value {
	thisObject := call.thisObject()
	kind := typedArrayKindFromClass(thisObject.class)

	length := int64(toUint32(thisObject.get("length")))
	byteLength := toUint32(thisObject.get("byteLength"))
	start, end := rangeStartEnd(call.ArgumentList, length, false)

	if start >= end {
		// Always an empty array
		buffer := call.runtime.newArrayBuffer(0)
		return toValue_object(call.runtime.newTypedArray(0, buffer, 0, byteLength, kind))
	}
	sliceLength := end - start
	sliceValueArray := make([]Value, sliceLength)

	for index := int64(0); index < sliceLength; index++ {
		from := arrayIndexToString(index + start)
		if thisObject.hasProperty(from) {
			sliceValueArray[index] = thisObject.get(from)
		}
	}

	buffer := call.runtime.newArrayBuffer(byteLength)
	typedArray := call.runtime.newTypedArray(uint32(sliceLength), buffer, 0, byteLength, kind)
	for i, value := range sliceValueArray {
		k := arrayIndexToString(int64(i))
		typedArray.put(k, value, true)
	}
	return toValue_object(typedArray)
}

func builtinTypedArray_some(call FunctionCall) Value {
	return builtinArray_some(call)
}

func builtinTypedArray_sort(call FunctionCall) Value {
	return builtinArray_sort(call)
}

func builtinTypedArray_subarray(call FunctionCall) Value {
	this := call.thisObject()
	buffer := this.get("buffer")._object()
	srcLength := this.get("length").number().int64
	relativeBegin := call.Argument(0).number().int64

	var beginIndex int64
	if relativeBegin < 0 {
		beginIndex = _maxInt64(srcLength+relativeBegin, 0)
	} else {
		beginIndex = _minInt64(relativeBegin, srcLength)
	}

	var relativeEnd int64
	end := call.Argument(1)
	if end.IsUndefined() {
		relativeEnd = srcLength
	} else {
		relativeEnd = end.number().int64
	}

	var endIndex int64
	if relativeEnd < 0 {
		endIndex = _maxInt64(srcLength+relativeEnd, 0)
	} else {
		endIndex = _minInt64(relativeEnd, srcLength)
	}

	newLength := _maxInt64(endIndex-beginIndex, 0)
	kind := typedArrayKindFromClass(this.class)
	elementSize := int64(typedArrayElementSize(kind))
	srcByteOffset := this.get("byteOffset").number().int64
	beginByteOffset := uint32(srcByteOffset + beginIndex*elementSize)
	byteLength := uint32(newLength * elementSize)
	typedArray := call.runtime.newTypedArray(uint32(newLength), buffer, beginByteOffset, byteLength, kind)
	return toValue_object(typedArray)
}

func builtinTypedArray_toLocaleString(call FunctionCall) Value {
	return builtinArray_toLocaleString(call)
}

func builtinTypedArray_toString(call FunctionCall) Value {
	return builtinArray_toString(call)
}

func builtinTypedArray_values(call FunctionCall) Value {
	panic("Unimplemented")
}
