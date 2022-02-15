package person

type person struct {
	firstName string
	lastName  string
	age       uint8
	adult     bool
}

func New(firstName, lastname string, age uint8, adult bool) *person {
	if age == 0 {
		age = 99
	}
	return &person{
		firstName: firstName,
		lastName:  lastname,
		age:       age,
		adult:     adult,
	}
}
