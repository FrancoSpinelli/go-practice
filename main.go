package main

import (
	"fmt"

	"github.com/FrancoSpinelli/go-practice/person"
)

func main() {
	franco := person.NewEmployed("Franco", "Spinelli", 23, true, 2, 100000)
	fmt.Printf("%+v \n", franco)
	fmt.Printf("%T \n", franco)
	franco.Person.Greet()
	franco.Human.Greet()
	franco.Payroll()
}
