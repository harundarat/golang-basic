package main

import "fmt"

func endApp() {
	fmt.Println("App ended")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Whopp.. Error occured")
	}
}

func main() {
	runApp(false)
	runApp(true)
}
