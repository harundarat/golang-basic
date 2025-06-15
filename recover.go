package main

import "fmt"

func endApp() {
	fmt.Println("App ended")
	message := recover()
	fmt.Println("Panic: ", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Whoopp... Error occured")
	}
}

func main() {
	runApp(true)
}
