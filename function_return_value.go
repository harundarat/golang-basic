package main

import "fmt"

func getHello(fullName string) string {
	hello := "Hello " + fullName
	return hello
}

func main() {
	result := getHello("Harun Al Rasyid")
	fmt.Println(result)

	fmt.Println(getHello("Harun"))
	fmt.Println(getHello("Al"))
	fmt.Println(getHello("Rasyid"))
}
