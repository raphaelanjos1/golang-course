package main

import (
	"fmt"
)

func main() {
	// declaração de variavel
	x := 10
	y := "olá"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	// atribuição de valor.
	x = 15
	fmt.Printf("x: %v, %T\n", x, x)
}