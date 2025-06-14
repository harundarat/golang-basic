package main

import "fmt"

func main() {
	person := make(map[string]string)
	person["name"] = "Harun Al Rasyid"
	person["country"] = "Indonesia"
	fmt.Println(person)
	fmt.Println(person["country"])

	location := map[string]string{
		"province": "Central Java",
		"street":   "wrong street",
	}

	fmt.Println(location)
	fmt.Println("len(location): ", len(location))
	delete(location, "street")
	fmt.Println(location)
}
