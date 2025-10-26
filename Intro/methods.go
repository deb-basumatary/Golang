package main

import "fmt"

// Add methods to types
type Vertex struct {
	x, y float32
}

// Receive argument
func (v Vertex) Area() float32 {
	return v.x * v.y
}
func main() {

	m := Vertex{4, 5}
	fmt.Println(m.Area())

}
