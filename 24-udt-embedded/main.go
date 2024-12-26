package main

import (
	"fmt"
)

func main() {

	var a1 struct {
		City    string
		PinCode string
	}

	a1.City = "Blr"
	a1.PinCode = "1223"

	fmt.Println(a1)
	var a2 struct {
		City    string
		PinCode string
	} = struct {
		City    string
		PinCode string
	}{
		City:    "Blr",
		PinCode: "312",
	}
	fmt.Println(a2)

	p1 := new(Person)
	p1.Name = "Jiten"
	p1.Email = "jitenp@outlook.com"
	p1.Address.City = "Blr"
	p1.Address.PinCode = "21232"

	p2 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Address: struct {
		City    string
		PinCode string
	}{
		City:    "Blr",
		PinCode: "1231",
	},
	}
	fmt.Println(p2)

}

type Person struct {
	Name, Email string
	Address     struct { // embedded struct
		City    string
		PinCode string
	}
}

func (p *Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("City:", p.Address.City)
	fmt.Println("PinCode:", p.Address.PinCode)
}

func (p *Person) DisplayAddress() {
	fmt.Println("City:", p.Address.City)
	fmt.Println("PinCode:", p.Address.PinCode)
}
