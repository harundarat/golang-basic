package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("Finish")

	// With Statement
	for count := 0; count <= 10; count++ {
		fmt.Println(count)
	}
	fmt.Println("Finish")

	// Range
	names := [...]string{"Harun", "Al", "Rasyid"}

	for index, value := range names {
		fmt.Println(index, ": ", value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	person := map[string]string{"name": "Harun Al Rasyid", "country": "Indonesia", "age": "22"}

	for key, value := range person {
		fmt.Println(key, ": ", value)
	}

}
