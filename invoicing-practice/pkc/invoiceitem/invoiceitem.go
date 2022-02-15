package invoiceitem

type invoiceitem struct {
	country string
	city    string
	total   float64
}

func New(country, city string, total float64) invoiceitem {
	return invoiceitem{country, city, total}
}
