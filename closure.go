package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("Counter: ", counter)
	increment()
	increment()
	increment()
	fmt.Println("Counter: ", counter)

}
