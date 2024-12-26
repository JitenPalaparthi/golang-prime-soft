package main

import "fmt"

func main() {
	finalData := New(100).Add(10).Add(20).Sub(10).Mul(5).Get()
	fmt.Println(finalData)

	var c Car
	c.SetColour("Red")
	c.SetFuelType("Petrol").SetTransmission("MT")
	fmt.Println(c)
}

func New(d int) *Calc {
	return &Calc{data: d}
}

type Calc struct {
	data int
}

func (c *Calc) Add(n int) ICalc {
	c.data += n
	return c
}
func (c *Calc) Sub(n int) ICalc {
	c.data -= n
	return c
}
func (c *Calc) Mul(n int) ICalc {
	c.data *= n
	return c
}

func (c *Calc) Get() int {
	return c.data
}

type ICalc interface {
	Add(n int) ICalc
	Sub(n int) ICalc
	Mul(n int) ICalc
	Get() int
}

type Car struct {
	colour      string
	fuelType    string
	tranmission string
}

func (c *Car) SetColour(clr string) ICar {
	c.colour = clr
	return c
}

func (c *Car) SetFuelType(ft string) ICar {
	c.fuelType = ft
	return c
}

func (c *Car) SetTransmission(t string) ICar {
	c.tranmission = t
	return c
}

type ICar interface {
	SetColour(c string) ICar
	SetFuelType(f string) ICar
	SetTransmission(f string) ICar
}
