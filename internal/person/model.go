package person

import "fmt"

type Person struct {
	Name   string
	Adress string
	Phone  string
}

func NewPerson(name, adress, phone string) *Person {
	return &Person{
		Name:   name,
		Adress: adress,
		Phone:  phone,
	}
}

func (p *Person) PrintAll() string {
	return fmt.Sprintf("Name   : %s\nAdress : %s\nPhone  : %s\n", p.Name, p.Adress, p.Phone)
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hello %s", p.Name)
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}
