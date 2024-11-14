package main

import (
	"CalculatorType/calc"
	"fmt"
)

func main() {
	var (
		a, b float64
		symb string
	)

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&symb)

	calc1 := calc.NewCalculator(a, symb, b)

	fmt.Println(calc1.Calculate(a, b, symb))
}