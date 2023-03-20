package main
import ("errors"
		"fmt"
		"log"
	)

func main()	{

	res, err := soma(10, 10)
	if err !=nil {
		log.fatal.(err.Error())
	}
	fmt.println(res)
}

func soma(x int, y int) (int, error) {
	res := x + y
	if res >10 {
		return 0, erros.New("Total maior que 10")
	}
}