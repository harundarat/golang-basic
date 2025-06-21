package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Banjarnegara", "Central Java", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Surakarta", "Central Java", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}
