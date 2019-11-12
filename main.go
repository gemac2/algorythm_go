package main

import (
	"fmt"
)

var a, b int = 5, 4

func main() {
	fmt.Println("Exercice 1")
	fmt.Println("Hello World!")
	fmt.Println("Exercice 2")
	fmt.Println(a, "+", b, "is", sum(a, b))
}

func sum(a, b int) int {
	result := a + b
	return result
}
