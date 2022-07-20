package main

import "fmt"

func main() {
	// booleans
	var flag bool // false by default
	var isAwesome = true
	fmt.Println(flag, isAwesome)

	// integers (signed and unsigned)
	var i int                           // 0 by default
	var j int8 = 127                    // range: -128 to 127
	var k int16 = 32767                 // range: -32768 to 32767
	var l int32 = 2147483647            // range: -2147483648 to 2147483647 (default ints are int32)
	var m int64 = 9223372036854775807   // range: -9223372036854775808 to 9223372036854775807 (on 64 bit machines, hold ints)
	var n uint = 0                      // range: 0 to 4294967295
	var o uint8 = 255                   // range: 0 to 255
	var p uint16 = 65535                // range: 0 to 65535
	var q uint32 = 4294967295           // range: 0 to 4294967295
	var r uint64 = 18446744073709551615 // range: 0 to 18446744073709551615
	var b byte = 255                    // range: 0 to 255 **bytes are same as uint8
	fmt.Println(i, j, k, l, m, n, o, p, q, r, b)

	// Explicit Type Casting/Conversion
	// Need to explicitly cast integer types in order to interact with each other
	var int_x int = 10
	var float64_y float64 = 30.2
	var addThemFloat float64 = float64(x) + y
	var addThemInt int = x + int(y)

	// floats
	var f float32 = 3.141592653589793 // default float32
	var g float64 = 3.141592653589793 // default float64
	fmt.Println(f, g)

	// complex numbers
	var c complex64 = 1 + 2i  // default complex64
	var d complex128 = 1 + 2i // default complex128
	fmt.Println(c, d)

	// strings
	var s string = "hello world"
	fmt.Println(s)

	// runes (unicode characters)
	var r1 rune = 'a'      // default rune
	var r2 rune = '\u0041' // unicode character
	fmt.Println(r1, r2)

	// arrays (fixed length)
	var a [2]string = [2]string{"hello", "world"}
	fmt.Println(a)

	// slices (variable length)
	var s1 []string = []string{"hello", "world"}
	var xx []int = []int{1, 2, 3, 4, 5}
	var xx2 []int = []int{6, 7, 8, 9, 10}
	var yy []int = []int{1, 4: 777, 9, 10: 666}
	var zz []int
	fmt.Println(s1, xx, yy)
	fmt.Println(zz)
	fmt.Println(len(xx), append(xx, 6))
	fmt.Println(len(xx2), append(xx, xx2...))

}
