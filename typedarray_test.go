package otto

import (
	"testing"
)

func TestTypedArray(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var a = new Int8Array(10)
            a.length
        `, 10)

		test(`
            var a = new Int16Array([0, 200, 0])
            a[1]
        `, 200)

		test(`
            var a = new Int8Array(4)
            a[0] = 300
            a[0]
        `, 44)

		test(`
            var a = new Uint8Array(1)
            a[0] = -100
            a[0]
        `, 156)

		test(`
            var a = new Uint8ClampedArray(1)
            a[0] = 260
            a[0]
        `, 255)

		test(`
            var a = new Uint8ClampedArray(1)
            a[0] = -100
            a[0]
        `, 0)

		test(`
            var a = new Float32Array(1)
            a[0] = 3.14159
            a[0]
        `, float32(3.14159))

		test(`
            var a = new Float64Array(2)
            a[1] = 1.617283904
            a[1]
        `, 1.617283904)
	})
}

func TestTypedArray_copyWithin(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            var a = new Int8Array(10)
            a[0] = 6
            a[1] = 5
            a[2] = 2
            a[3] = 10

            a.copyWithin(4, 1)
            a[4]
        `, 5)
	})
}

func TestTypedArray_every(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array(2)
			a[0] = 1
			a[1] = 2
			a.every(function(x) { return x !== 0 })
		`, true)

		test(`
			var a = new Int8Array(2)
			a[0] = 1
			a[1] = 2
			a.every(function(x) { return x === 1 })
		`, false)
	})
}

func TestTypedArray_fill(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array(5)
			a.fill(10)
		`, "10,10,10,10,10")

		test(`
			var a = new Int8Array(5)
			a.fill(10, 2)
		`, "0,0,10,10,10")

		test(`
			var a = new Int8Array(5)
			a.fill(10, 2, 4)
		`, "0,0,10,10,0")

		test(`
			var a = new Int8Array(5)
			a.fill(128)
		`, "-128,-128,-128,-128,-128")
	})
}

func TestTypedArray_filter(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Uint8Array(10)
			a[3] = 1
			a[5] = 1
			a[7] = 1
			a.filter(function(x) { return x === 1 })
		`, "1,1,1")
	})
}

func TestTypedArray_find(t *testing.T) {
	tt(t, func() {
		test, _ := test()
		test(`
			var a = new Int8Array([0, 2, 0])
			a.find(function(x) { return x === 2 })
		`, 2)

		test(`
			var a = new Int32Array(0)
			a.find(function(x) { return x === 1 })
		`, UndefinedValue())
	})
}

func TestTypedArray_findIndex(t *testing.T) {
	tt(t, func() {
		test, _ := test()
		test(`
			var a = new Int8Array([0, 2, 0])
			a.findIndex(function(x) { return x === 2 })
		`, 1)

		test(`
			var a = new Int32Array(0)
			a.findIndex(function(x) { return x === 1 })
		`, -1)
	})
}

func TestTypedArray_forEach(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int32Array([10, 20, 30])
			var b = []
			a.forEach(function(x) { b.push(x) })
			b
		`, "10,20,30")
	})
}

func TestTypedArray_includes(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int32Array([0, 10, 40, 2])
			a.includes(2)
		`, true)

		test(`
			var a = new Int32Array([0, 10, 40, 2])
			a.includes(1)
		`, false)
	})
}

func TestTypedArray_indexOf(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int32Array([10, 100, 1000])
			a.indexOf(100)
		`, 1)

		test(`
			var a = new Int32Array([10, 100, 1000])
			a.indexOf(10000)
		`, -1)
	})
}

func TestTypedArray_join(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int32Array([10, 100, 1000])
			a.join(";")
		`, "10;100;1000")
	})
}

func TestTypedArray_lastIndexOf(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int32Array([10, 20, 10, 30])
			a.lastIndexOf(10)
		`, 2)

		test(`
			var a = new Int32Array([10, 20, 10, 30])
			a.lastIndexOf(1)
		`, -1)
	})
}

func TestTypedArray_map(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([2, 4, 8])
			a.map(function(x) { return x * 2 })
		`, "4,8,16")

		test(`
			var a = new Int8Array([64, 72, 4])
			a.map(function(x) { return x * 2 })
		`, "-128,-112,8")
	})
}

func TestTypedArray_reduce(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([2, 4, 8])
			a.reduce(function(sum, x) { return sum + x }, 0)
		`, 14)
	})
}

func TestTypedArray_reduceRight(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([2, 4, 8])
			a.reduce(function(sum, x) { return sum + x }, 0)
		`, 14)
	})
}

func TestTypedArray_reverse(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([2, 4, 8])
			a.reverse()
		`, "8,4,2")
	})
}

func TestTypedArray_set(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array(3)
			a.set([1,2,3])
			a
		`, "1,2,3")

		test(`
			var a = new Int8Array(3)
			var b = new Int32Array([128, 64, 32])
			a.set(b)
			a
		`, "-128,64,32")
	})
}

func TestTypedArray_slice(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([1, 2, 3, 4, 5])
			a.slice(2)
		`, "3,4,5")
	})
}

func TestTypedArray_some(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([1, 0, 1])
			a.some(function(x) { return x === 1 })
		`, true)

		test(`
			var a = new Int8Array([0, 0, 0])
			a.some(function(x) { return x === 1 })
		`, false)
	})
}

func TestTypedArray_sort(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int8Array([4, 2, 8, 6])
			a.sort()
			a
		`, "2,4,6,8")
	})
}

func TestTypedArray_subarray(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
			var a = new Int16Array([1, 2, 3, 4])
			var b = a.subarray(2)
			b
		`, "3,4")
	})
}

func TestTypedArray_toLocaleString(t *testing.T) {
	tt(t, func() {
		test, _ := test()
		test(`
			var a = new Int8Array([4,5,6])
			a.toLocaleString()
		`, "4,5,6")
	})
}

func TestTypedArray_toString(t *testing.T) {
	tt(t, func() {
		test, _ := test()
		test(`
			var a = new Int8Array([4,5,6])
			a.toString()
		`, "4,5,6")
	})
}
