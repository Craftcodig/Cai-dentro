package 

import "fmt"

func main() {

	variavel := 10
	
	abc(&variavel)
	
	fmt.Println(variavel)

}

func abc(a *int) {
	
	*a= 200
}