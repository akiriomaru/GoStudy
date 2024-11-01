package maths

import "fmt"

func Sumresult(a float32, b float32) float32 {
	return a+b
}

func Difresult(a float32, b float32) float32 {
	return a-b
}

func Multresult(a float32, b float32) float32 {
	return a*b
}

func Divresult(a float32, b float32) float32 {
	if b != 0 {return a/b}
	fmt.Println("Попытка деления на ноль!")
	return 0
}
