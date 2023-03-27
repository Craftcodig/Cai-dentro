package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main(){

	contador("sem go routine")
	go contador("com go routine")

	fmt.Println("Produto 1")
	fmt.Println("Produto 2")
	fmt.Println("fim...")
	time.Sleep(time.Second)

}
