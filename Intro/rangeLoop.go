package main

import (
	"fmt"
)

func main() {
	// Alternative way
	// var arr [6]int = [6]int{1, 2, 3, 4, 3, 5}
	//Idiomatic
	arr := [6]int{4, 3, 5, 1, 3, 6}

	// range over an array or slice
	for i, v := range arr {
		fmt.Println("index: ", i, "value: ", v)
	}
	// discard index
	for _, v := range arr {
		fmt.Println(v)
	}
	// Only Index
	for i := range arr {
		fmt.Println(i)
	}
	fmt.Println(arr)
}
