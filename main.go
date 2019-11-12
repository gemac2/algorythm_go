package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercices 1")
	fmt.Println("Hello World!")
	fmt.Println("Exercices 2")
	fmt.Println(" 5 + 4 is", sum(5, 4))
}

func sum(a int, b int) int {
	result := a + b
	return result
}
