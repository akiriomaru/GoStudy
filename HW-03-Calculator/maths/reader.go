package maths

import "fmt"

var (
	a, b float32
	symb string
)

func Readexpression() (float32, string, float32) {
    fmt.Scanln(&a)
	fmt.Scanln(&symb)
	fmt.Scanln(&b)
	return a, symb, b
}
