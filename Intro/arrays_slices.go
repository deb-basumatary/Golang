package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"
	fmt.Println(a[0], " ", a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	even := [...]int{2, 4, 6, 8, 10, 12}
	fmt.Println(even)
	//2D arrray
	matrix := [2][3]int{
		{2, 5, 3},
		{7, 6, 4},
	}
	fmt.Println(matrix)

	// Slice aka Go's Dynamic Arrays
	s := primes[1:3] // [start, end)
	fmt.Println(s)
	fmt.Printf("Length: %v and Current Capacity: %v \n", len(s), cap(s))

	// append value, need to pass the slice and also accept the output
	s = append(s, 17, 19, 23)
	fmt.Println(s)

}
