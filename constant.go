package main

import "fmt"

func main() {
	const firstName string = "Harun"
	const lastName = "Al Rasyid"
	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		age     = 22
		country = "Indonesia"
	)
	fmt.Println(age)
	fmt.Println(country)
}
