package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo = contactInfo{email: "omarghetti2@gmail.com", zipCode: 29027}
	alex.setName("Omar")
	alex.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) setName(newFirstName string) {
	p.firstName = newFirstName
}
