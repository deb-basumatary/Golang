package main

import (
	"fmt"
	"math"
)

func power(x, n, lim float64) float64 {
	// Like for if can start with a short statement before the condition
	// Variable declared by the statement are only scope until the end of the if and else block
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		power(3, 2, 10),
		power(3, 3, 20),
	)
}
