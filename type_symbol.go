package otto

type _symbol struct {
	description string
}

// func (runtime *_runtime) newSymbolObject() *_object {
// 	self := runtime.newObject()
// 	self.class = "Symbol"
// 	self.objectClass = _classObject
// 	self.value = _symbol{}
// 	self.defineProperty("length", toValue_uint32(0), 0100, false)
// 	return self
// }

func (runtime *_runtime) newSymbolObject(value Value) *_object {
	return runtime.newPrimitiveObject("Symbol", toValue_symbol(value._symbol()))
}

func (self *_object) symbolValue() _symbol {
	value, _ := self.value.(Value)
	return value._symbol()
}
