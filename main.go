package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercices 1 \n"+"Hello World! \n"+"Exercices 2 \n", sum(5, 4))
}

func sum(a int, b int) int {
	var c int = a + b
	return c
}
