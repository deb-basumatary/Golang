package main

import (
	"fmt"
	"math"
)

// constant Value
const PI = 3.14

// data type comes after variable
func add(x int, y int) int {
	return x + y
}

// If consecutive variables have same data types, we can add the data type after the last variables
func multiply(x, y int) int {
	return x * y
}

func mixedDT(x, y int, name string) string {
	return "Hello"
}

// Return multiple values
func valueReturn() (string, string) {
	return "Hello", "World!"
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("Type of Z: %T", z)
}

func main() {
	fmt.Println(add(2, 34))
	fmt.Println(multiply(30, 4))
	fmt.Println(mixedDT(2, 4, "Hola"))

	// var used to declare variable, consecutive variables with same data types can be grouped.
	var h, w string = valueReturn()
	fmt.Println("H value:", h)
	fmt.Println("W value:", w)
	typeConversion()
}
