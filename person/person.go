package person

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       uint8
	adult     bool
}

// New constructor to person
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

func (p *person) SetFirstName(firstName string) { p.firstName = firstName }
func (p *person) FirstName() string             { return p.firstName }

func (p *person) SetLastName(lastName string) { p.lastName = lastName }
func (p *person) LastName() string            { return p.lastName }

func (p *person) SetAge(age uint8) { p.age = age }
func (p *person) Age() uint8       { return p.age }

func (p *person) SetAdult(adult bool) { p.adult = adult }
func (p *person) Adult() string       { return p.firstName }

// Greet to person
func (p *person) Greet() {
	fmt.Printf("Welcome back %s %s", p.firstName, p.lastName)
}
