package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 *Address = new(Address)
	var address2 *Address = address1

	address2.City = "Banjarnegara"

	fmt.Println(address1)
	fmt.Print(address2)
}
