package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "jitenp@outlook.com", Address: &Address{City: "Blr", Pincode: "560086"}}
	p2 := new(Person)
	p2.Name = "Jiten"
	p2.Email = "jitenp@outlook.com"
	p2.Address.City = "Blr"
	p2.Address.Pincode = "560086"
	fmt.Println(p1)
	fmt.Println(p2)

}

type Person struct {
	Name, Email string
	Address     *Address // composion
}

type Address struct {
	City, Pincode string
}
