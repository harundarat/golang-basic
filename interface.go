package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name, Country string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	harun := Person{"Harun Al Rasyid", "Indonesia"}

	SayHello(harun)
}
