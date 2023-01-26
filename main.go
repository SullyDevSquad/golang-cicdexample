package main

import (
	"fmt"
	"golang-cicdexample/calculator"
)

func main() {
	fmt.Println("Hello, calculator package!")
	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Subtract(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
	fmt.Println(calculator.Divide(1, 2))
	fmt.Println("Program finished")
}
