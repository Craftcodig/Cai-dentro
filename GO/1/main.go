package main

import "fmt"

// T1
func main()
{
	// T1 <-> T2
produto := make(chan string)

	// T2
go func() {
	produto <- "Produto Disponível"
}()

result := <-produto
fmt.Println(result)
}