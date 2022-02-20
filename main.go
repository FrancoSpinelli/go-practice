package main

import (
	"fmt"

	"github.com/FrancoSpinelli/go-practice/person"
)

func main() {
	franco := person.NewEmployed("Franco", "Spinelli", 23, true, 100000)
	fmt.Printf("%+v \n", franco)
	fmt.Printf("%T \n", franco)
	franco.Greet()
	franco.Payroll()
}
