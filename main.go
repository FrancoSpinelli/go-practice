package main

import "fmt"

type Person struct {
	Id   uint
	Name string
	Age  uint8
	Live bool
}

func (p *Person) greet() {
	fmt.Printf("Welcome back %s\n", p.Name)
}

func (p *Person) changeName(name string) {
	p.Name = name
}

func main() {
	Franco := Person{
		5,
		"Franco Spinelli",
		23,
		true,
	}
	Franco.greet()
	Franco.changeName("Max Power")
	Franco.greet()

}
