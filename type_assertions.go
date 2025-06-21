package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()

	// resultString := result.(string)
	// fmt.Println("resultString: ", resultString)

	// // resultInt := result.(int) // panic
	// // fmt.Println("resultInt: ", resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String: ", value)
	case int:
		fmt.Println("Int: ", value)
	default:
		fmt.Println("Unknown: ", value)
	}
}
