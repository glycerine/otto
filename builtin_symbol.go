package otto

import "fmt"

func builtinSymbol(call FunctionCall) Value {
	description := call.Argument(0)
	var sym string
	if description.IsDefined() {
		sym = description.string()
	} else {
		sym = ""
	}
	return toValue_symbol(_symbol{sym})
}

func builtinNewSymbol(self *_object, argumentList []Value) Value {
	panic(self.runtime.panicTypeError("Symbol is not a constructor"))
}

func builtinSymbol_toString(call FunctionCall) Value {
	this := call.thisObject()
	str := fmt.Sprintf("Symbol(%s)", this.symbolValue().description)
	return toValue_string(str)
}

func builtinSymbol_valueOf(call FunctionCall) Value {
	return call.This
}

func builtinSymbol_for(call FunctionCall) Value {
	str := call.Argument(0).string()
	if sym, ok := call.runtime.symbolRegistry[str]; ok {
		return toValue_symbol(sym)
	}

	sym := _symbol{str}
	call.runtime.symbolRegistry[str] = sym
	return toValue_symbol(sym)
}

func builtinSymbol_keyFor(call FunctionCall) Value {
	arg := call.Argument(0)
	if !arg.IsSymbol() {
		panic(call.runtime.panicTypeError(fmt.Sprintf("%s is not a symbol", arg.string())))
	}

	return toValue_string(arg._object().symbolValue().description)
}
