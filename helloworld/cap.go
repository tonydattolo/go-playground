package main

import "fmt"

func main() {
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))
	// NOTE: notice how the capacity of the array grows as the array is filled
	// 		 it uses a doubling function as needed.

	// make allows you to specify type, length, and capacity of the slice
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))
	// if you were to call append here, it would be placed after the 5th element and 0s for [0]-[4]
	// it would also double the capacity of the slice to 10, because append always grows, and the inital
	// values were prefilled with 0s

	x = append(x, 10)
	fmt.Printf("x: %d, len: %d, cap: %d\n", x, len(x), cap(x))
	// x2 := make([]int, 5, 10)

	// empty slice literal
	var y = []int{} // 0 len non-nil slice
	fmt.Println(y, len(y), cap(y))

	data := []int{2, 4, 6, 8, 10}
	fmt.Println(data, len(data), cap(data))
}
