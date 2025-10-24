package main

import (
	"fmt"

	"rsc.io/quote" // External module, downloaded after using go mod tidy
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(quote.Go())

}
