package calc

import "fmt"

type calculator struct {
	first, second float64
	symbol string
}

func NewCalculator(a float64, symb string, b float64) calculator {
	return calculator{
		first: a,
		second: b,
		symbol: symb,
	}
}

func (c calculator) Calculate(a, b float64, symb string) float64 {
	switch symb {
		case "+": return sumresult(a, b)
		case "-": return difresult(a, b)
		case "*": return multresult(a, b)
		case "/": return divresult(a, b)
		default: fmt.Println("Unknown Number")
		}
	return 0
}

func sumresult(a float64, b float64) float64 {
	return a+b
}

func difresult(a float64, b float64) float64 {
	return a-b
}

func multresult(a float64, b float64) float64 {
	return a*b
}

func divresult(a float64, b float64) float64 {
	if b != 0 {return a/b}
	fmt.Println("Попытка деления на ноль!")
	return 0
}