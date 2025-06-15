package main

import "fmt"

func logging() {
	fmt.Println("Function call finished")
}

func runApplication() {
	defer logging()

	fmt.Println("Run application")
}

func main() {
	runApplication()
}
