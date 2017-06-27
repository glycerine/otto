package otto

func builtinArrayBuffer(call FunctionCall) Value {
	return toValue_object(builtinNewArrayBufferNative(call.runtime, call.ArgumentList))
}

func builtinNewArrayBuffer(self *_object, argumentList []Value) Value {
	return toValue_object(builtinNewArrayBufferNative(self.runtime, argumentList))
}

func builtinNewArrayBufferNative(runtime *_runtime, argumentList []Value) *_object {
	return runtime.newArrayBuffer(arrayUint32(runtime, argumentList[0]))
}

func builtinArrayBuffer_slice(call FunctionCall) Value {
	arrayBuffer := arrayBufferObjectOf(call.runtime, call.thisObject())
	length := int64(len(arrayBuffer.buffer))
	start, end := rangeStartEnd(call.ArgumentList, length, false)

	if start >= end {
		// Always an empty array buffer
		return toValue_object(call.runtime.newArrayBuffer(0))
	}

	sliceLength := end - start
	slice := make([]byte, sliceLength)
	copy(slice, arrayBuffer.buffer[start:end])
	return toValue_object(call.runtime.newArrayBufferObjectOf(slice))
}

func builtinArrayBuffer_isView(call FunctionCall) Value {
	return toValue_bool(call.Argument(0).IsTypedArray())
}
