package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// This is not using pointer => address2 will copy and create new data from address1
	// address1 := Address{"Banjarnegara", "Central Java", "Indonesia"}
	// address2 := address1

	// address2.City = "Solo"

	// fmt.Println(`Address1: `, address1) // Address1:  {Banjarnegara Central Java Indonesia}
	// fmt.Println(`Address2: `, address2) // Address2:  {Solo Central Java Indonesia}

	// This is using pointer => address2 will be only as reference to address1
	address1 := Address{"Banjarnegara", "Central Java", "Indonesia"}
	address2 := &address1

	address2.City = "Solo"
	fmt.Println(`Address1: `, address1) // Address1:  {Solo Central Java Indonesia}
	fmt.Println(`Address2: `, address2) // Address2:  &{Solo Central Java Indonesia}

}
