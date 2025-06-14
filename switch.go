package main

import "fmt"

func main() {
	name := "Rasyid"

	switch name {
	case "Harun":
		fmt.Println("Hello Harun")
	case "Al":
		fmt.Println("Hello Al")
	case "Rasyid":
		fmt.Println("Hello Rasyid")
	default:
		fmt.Println("Hi, What's your name?")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name too long")
	case false:
		fmt.Println("Name OK")
	}

	// No expression
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name too long")
	default:
		fmt.Println("Name OK")
	}
}
