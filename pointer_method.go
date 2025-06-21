package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	harun := Man{Name: "Harun"}

	harun.Married()

	fmt.Println(harun.Name)
}
