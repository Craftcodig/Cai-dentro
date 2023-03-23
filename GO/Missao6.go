packpage main

import "fmt"

func main() {
	// memoria-0xc000160a0(10) <------- A <------- 10

	a :=10
	var ponteiro *int = &a
	*ponteiro = 50
	b :=&a
	*b =60
	fmt.Println(a)
	
}