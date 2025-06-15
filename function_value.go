package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	example := getGoodBye
	bye := getGoodBye

	fmt.Println(example("Harun"))
	fmt.Println(bye("Rasyid"))
}
