package main

import "fmt"

func main() {
	resultado := Soma(1, 1)
	fmt.Printf("%v", resultado)
}

func Soma(a int, b int) int {
	return a + b
}