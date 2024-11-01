package main

import (
	"HW-03-Calculator/maths"
	"fmt"
)

var (
	a, b float32
	symb string
)

func main() {
	a, symb, b = maths.Readexpression()
	switch symb {
	case "+": fmt.Println("Результат сложения: ", maths.Sumresult(a, b))
	case "-": fmt.Println("Результат вычитания: ", maths.Difresult(a, b))
	case "*": fmt.Println("Результат умножения: ", maths.Multresult(a, b))
	case "/": fmt.Println("Результат деления: ", maths.Divresult(a, b))
	default: fmt.Println("Unknown Number")
	}
}
