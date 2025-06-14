package main

import "fmt"

// func getCompleteName() (firstName string, middleName string, lastName string) {
// 	firstName = "Harun"
// 	middleName = "Al"
// 	lastName = "Rasyid"

// 	return firstName, middleName, lastName
// }

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Harun"
	middleName = "Al"
	lastName = "Rasyid"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}
