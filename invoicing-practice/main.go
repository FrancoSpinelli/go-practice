package main

import (
	"fmt"

	"github.com/FrancoSpinelli/go-practice/invoicing-practice/pkg/customer"
	"github.com/FrancoSpinelli/go-practice/invoicing-practice/pkg/invoice"
	"github.com/FrancoSpinelli/go-practice/invoicing-practice/pkg/invoiceitem"
)

func main() {
	item1 := invoice.New(
		"Argentina",
		"Capital",
		customer.New("Franco", "Av. Cordoba", "1555896981"),
		invoiceitem.NewItem(
			invoiceitem.New(1, "Yerba", 5.89),
			invoiceitem.New(2, "Mate", 7.98),
			invoiceitem.New(3, "Bombilla", 1.66),
		),
	)
	item1.SetTotal()
	fmt.Printf("%+v", item1)
}
