package main

import "net/http"

func main() {
	
	res, err := http.Get("http://google.com.br")
	uf err!= nil {
	log.Fatal(err.error())
}

	fmt.Println(res.Header)