package main

import "fmt"

func main() {

	// For loop
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	// For loop without init statement
	j := 0
	for ; j < 10; j++ {
		sum += j
	}
	// For Loop without post statement
	for j < 20 {
		j++
		sum += j
	}
	// while loop
	for sum < 1000 {
		sum += 100
	}
	fmt.Println(sum)
}
