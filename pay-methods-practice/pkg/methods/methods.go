package methods

type PayMethod interface {
	Pay() string
}

type creditCard struct{}
type cash struct{}
type paypal struct{}

func (c creditCard) Pay() string {
	return "Pay with credit card"
}

func (c cash) Pay() string {
	return "Pay with Cash"
}

func (c paypal) Pay() string {
	return "Pay with Paypal"
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return creditCard{}
	case 2:
		return cash{}
	case 3:
		return paypal{}
	default:
		return nil
	}
}
