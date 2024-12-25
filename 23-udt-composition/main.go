package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "jitenp@outlook.com", Address: &Address{City: "Blr", Pincode: "560086"}}
	p2 := new(Person)
	p2.Name = "Jiten"
	p2.Email = "jitenp@outlook.com"
	a2 := &Address{
		City:    "Blr",
		Pincode: "560086",
	}
	p2.Address = a2

	fmt.Println(p1)
	fmt.Println(p2)
	p1.Address.Display()

	s1 := Student{Name: "Jiten", Email: "jitenp@outlook.com", Status: "active", Address: Address{City: "Blr", Pincode: "560086", Status: "active"}}

	s1.Display()
	fmt.Println("Status of address", s1.Address.Status)

}

func New() *Person {

	return nil
}

type Person struct {
	Name, Email string
	Address     *Address // composion
}

type Address struct {
	City, Pincode, Status string
}

func (a *Address) Display() {
	fmt.Println("City:", a.City)
	fmt.Println("PinCode:", a.Pincode)
	fmt.Println("Status:", a.Status)
}

type Student struct {
	Name         string
	Email        string
	Status       string
	LastModified int64
	Address      // promoted field
}

// func (s *Student) GetName() string {
// 	return s.Name
// }

// func (s *Student) SetName(name string) {
// 	s.Name = name
// }
