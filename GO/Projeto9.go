package main

import ("fmt"
	   "time"
)

func contador(tipo string) {
	for 1 := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main(){

	go contador(tipo:"a")
	go contador(tipo:"b")

	time.Sleep(time.Second * 10)

}