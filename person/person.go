package person

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       uint8
	adult     bool
}

// New constructor to Person
func NewPerson(firstName, lastname string, age uint8, adult bool) *Person {
	if age == 0 {
		age = 99
	}
	return &Person{
		firstName: firstName,
		lastName:  lastname,
		age:       age,
		adult:     adult,
	}
}

func (p *Person) SetFirstName(firstName string) { p.firstName = firstName }
func (p *Person) FirstName() string             { return p.firstName }

func (p *Person) SetLastName(lastName string) { p.lastName = lastName }
func (p *Person) LastName() string            { return p.lastName }

func (p *Person) SetAge(age uint8) { p.age = age }
func (p *Person) Age() uint8       { return p.age }

func (p *Person) SetAdult(adult bool) { p.adult = adult }
func (p *Person) Adult() string       { return p.firstName }

// Greet to Person
func (p *Person) Greet() {
	fmt.Printf("Welcome back %s %s \n", p.firstName, p.lastName)
}

type Human struct {
	age      uint8
	children uint8
}

func NewHuman(age, children uint8) Human {
	return Human{age, children}
}

func (h Human) Greet() {
	fmt.Println("Welcome human")
}

type Employee struct {
	// parecido a una herencia múltiple
	Person
	Human
	salary float64
}

func NewEmployed(firstName, lastname string, age uint8, adult bool, children uint8, salary float64) Employee {
	return Employee{*NewPerson(firstName, lastname, age, adult), NewHuman(age, children), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.salary * 0.9)
}
