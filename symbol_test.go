package otto

import "testing"

func TestSymbol(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
		    var s = Symbol('foo')
		    s
		`, _symbol{"foo"})

		test(`
            var s = Symbol('foo')
            s.toString()
        `, "Symbol(foo)")
	})
}
