package main

import "fmt"

type Person struct {
	name    string
	state   string
	country string
}

func (p *Person) setName(name string) *Person {
	fmt.Println("in setName")
	fmt.Println(p)
	p.name = name
	return p
}
func (p *Person) setState(state string) *Person {
	fmt.Println("in setState")
	fmt.Println(p)
	p.state = state
	return p
}
func (p *Person) setCountry(country string) *Person {
	fmt.Println("in setCountry")
	fmt.Println(p)
	p.country = country
	return p
}

func main() {
	joe := Person{}
	fmt.Println(joe)
	joe.setName("joe").setState("MO").setCountry("USA")
	fmt.Println(joe)
}
