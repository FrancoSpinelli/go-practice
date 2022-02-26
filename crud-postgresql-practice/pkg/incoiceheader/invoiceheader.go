package invoiceheader

import "time"

// Model of InvoiceHeader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
