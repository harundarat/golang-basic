package main

import "fmt"

func main() {
	var age uint8
	age = 10
	fmt.Println(age)

	name := "Harun Al Rasyid"
	fmt.Println(name)

	var (
		firstName  = "Harun"
		middleName = "Al"
		lastName   = "Rasyid"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
