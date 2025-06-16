package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, " my name is ", customer.Name)
}

func main() {
	var harun Customer

	harun.Name = "Harun"
	harun.Address = "Indonesia"
	harun.Age = 22

	fmt.Println(harun.Name)
	fmt.Println(harun.Age)
	fmt.Println(harun)

	al := Customer{
		Name:    "Al",
		Address: "Indonesia",
		Age:     22,
	}
	fmt.Println(al)

	rasyid := Customer{"Rasyid", "Indonesia", 22}
	fmt.Println(rasyid)

	harun.sayHello("Vincent")
	al.sayHello("Vincent")
	rasyid.sayHello("Vincent")
}
