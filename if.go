package main

import "fmt"

func main() {
	name := "Vincent"

	if name == "Harun" {
		fmt.Println("Hello Harun")
	} else if name == "Al" {
		fmt.Println("Hello Al")
	} else if name == "Rasyid" {
		fmt.Println("Hello Rasyid")
	} else {
		fmt.Println("Hi, What's your name?")
	}

	// Short Statement

	if length := len(name); length > 5 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Name OK")
	}
}
