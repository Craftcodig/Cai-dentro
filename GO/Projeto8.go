package main

import "fmt"

func contador(tipo string) {
	for 1 := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main(){

	contador(tipo:"sem go routine")
	go contador(tipo:"com go routine")

	fmt.Println(a...:"Produto 1")
	fmt.Println(a...:"Produto 2")
	fmt.Println(a...:"fim...")
	time.Sleep(time.Second)

}
