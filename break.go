package main

import "fmt"

func main() {
	for count := 0; count <= 10; count++ {
		if count == 5 {
			break
		}

		fmt.Println("count: ", count)
	}
}
