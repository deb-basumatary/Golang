package main

import "fmt"

type rectangle struct {
	length, breadth int8
}

func main() {

	pi := 3.14

	pointer := &pi

	*pointer = 200

	rect := rectangle{length: 100, breadth: 45}

	pointer2 := &rect

	// In struct no need to use to * to access the value
	pointer2.breadth = 32

	fmt.Println(pointer2.length)
	fmt.Println((*pointer2).length)

}
