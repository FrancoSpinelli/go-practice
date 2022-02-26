package invoiceitem

import "time"

//Model of IinvoiceItem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdeatedAt      time.Time
}
