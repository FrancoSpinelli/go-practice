package invoiceitem

// Item is the structure of Item
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.Value
func (i Item) Value() float64 {
	return i.value
}
