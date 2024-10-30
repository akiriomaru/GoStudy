package main

import "fmt"

func main() {

	const (
		jan = iota + 1
		feb
		mar
		apr
	)

	fmt.Println(jan, feb, mar, apr)
}