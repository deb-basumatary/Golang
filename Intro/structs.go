package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 32
	fmt.Println(v.X)

	//Access fields through pointer
	p := &v
	p.X = 1.234e9
	fmt.Println(p.X)
	fmt.Println((*p).X)

	//Implicit value are used if none provided
	v2 := Vertex{X: 3}
	v3 := Vertex{}
	v4 := &Vertex{1, 3}
	fmt.Println(v2, v3, *v4)

}
