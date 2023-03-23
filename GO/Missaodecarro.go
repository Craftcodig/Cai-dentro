package main
import "fmt"
type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {
	carro := Carro{
		Name "Ka",
	}
	carro.andou()
	fmt.Println(carro.Name)
}