package main

import (
	"fmt"
)

func main() {
	i := 32
	p := &i
	fmt.Println("P value: ", p)
	fmt.Println("*P value: ", *p)
	//Changing i value through pointer
	*p = 32
	fmt.Println("I value: ", i)

	q := &p
	fmt.Println("Point P address: ", q)

	var p1 *string
	var name string = "Hello World"
	p1 = &name
	fmt.Printf("\np1 value: %v and address: %v", *p1, p1)
}
