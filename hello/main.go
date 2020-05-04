package main

import (
	"fmt"
)

func main() {
	x := 10
	y := "string"
	z := true
	
	numberofbytes, errors := fmt.Println("Hello World!")
	fmt.Println(numberofbytes, errors)
	fmt.Println(x, y, z)
}