package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "David"
	m[2] = "Das"

	n := map[int]string{1: "Jane", 2: "Street"}
	fmt.Println(m)
	fmt.Println(n)
}
