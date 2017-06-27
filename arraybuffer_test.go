package otto

import (
	"testing"
)

func TestArrayBuffer(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var buf = new ArrayBuffer(10)
            buf.byteLength
        `, 10)
	})
}

func TestArrayBuffer_slice(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var buf1 = new ArrayBuffer(32)
			var buf2 = buf1.slice(4)
			buf2.byteLength
		`, 28)

		test(`
			var buf1 = new ArrayBuffer(32)
			var buf2 = buf1.slice(8, 16)
			buf2.byteLength
		`, 8)
	})
}

func TestArrayBuffer_isView(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var buf = new ArrayBuffer(10)
            ArrayBuffer.isView(buf)
        `, false)

		test(`
			ArrayBuffer.isView([])
		`, false)

		test(`
			ArrayBuffer.isView(true)
		`, false)
	})
}
