package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Começou")
	time.Sleep(time.Second)
	fmt.Println("Terminou")
	}