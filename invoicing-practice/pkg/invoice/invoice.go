package invoice

import (
	"github.com/FrancoSpinelli/go-practice/invoicing-practice/pkg/customer"
	"github.com/FrancoSpinelli/go-practice/invoicing-practice/pkg/invoiceitem"
)

// Invoice is the structure of Invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

//New returns a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.Total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
