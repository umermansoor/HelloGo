package main

import "fmt"

type Person struct {
	name string
	year_of_birth uint
}

func (p *Person) get_name() string {
	return p.name
}


func main() {
	p := new(Person)
	p.name = "Umer"
	fmt.Println(p.get_name())

}
