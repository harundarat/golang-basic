package main

import "fmt"

func getFullName() (string, string) {
	return "Harun", "Al Rasyid"
}

func main() {
	// firstName, lastName := getFullName()
	firstName, _ := getFullName()

	fmt.Println(firstName)
	// fmt.Println(lastName)
}
