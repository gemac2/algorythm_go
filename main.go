package main

import (
	"fmt"
)

var a, b int = 5, 4
var number int = 0

func main() {
	fmt.Println("Exercice 1 - Hello World")
	fmt.Println("Hello World!")
	fmt.Println("Exercice 2 - Add two numbers")
	fmt.Println(a, "+", b, "is", sum(a, b))
	fmt.Println("Exercice 3 - Print numbers from 1 to 10")
	for i := 1; i <= 10; i++ {
		number = i
		fmt.Print(number, " ")
	}

}

func sum(a, b int) int {
	result := a + b
	return result
}
