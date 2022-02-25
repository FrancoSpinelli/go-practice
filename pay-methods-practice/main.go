package main

import (
	"fmt"

	"github.com/FrancoSpinelli/go-practice/pay-methods/pkg/methods"
)

func main() {
	var selectedMethod uint
	fmt.Println("Select pay method:")
	fmt.Println("\t 1: Credit Card 2: Cash 3: Paypal")

	_, err := fmt.Scanln(&selectedMethod)

	if err != nil {
		panic("Method invalid. Select a valid method")
	}
	if selectedMethod > 3 {
		panic("Method invalid. Select a valid method")
	}

	payMethod := methods.Factory(selectedMethod)
	fmt.Println(payMethod.Pay())
}
