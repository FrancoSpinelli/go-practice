package main

import (
	"fmt"

	"github.com/FrancoSpinelli/go-practice/person"
)

func main() {

	franco := person.New("Franco", "Spinelli", 23, true)
	fmt.Printf("%+v \n", franco)

}
