package main

import "fmt"

type person struct {
	fistName string
	lastName string
	contact  contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p := person{
		fistName: "Alex",
		lastName: "Party",
		contact: contactInfo{
			email:   "test@gmail.com",
			zipCode: 1000,
		},
	}

	//fmt.Printf("%p", &p)

	p.updateName("LL")
	p.print()
}

func (personPtr *person) updateName(name string) {
	(*personPtr).fistName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
