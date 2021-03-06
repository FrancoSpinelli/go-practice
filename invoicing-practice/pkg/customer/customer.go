package customer

// Customer is the structure of Customer
type Customer struct {
	name    string
	address string
	phone   string
}

// New returns a new Customer
func New(name, address string, phone string) Customer {
	return Customer{name, address, phone}
}
