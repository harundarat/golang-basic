package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Harun"
	names[1] = "Al"
	names[2] = "Rasyid"
	fmt.Println(names)

	ages := [3]uint8{22, 23}
	fmt.Println(ages)

	scores := [...]uint8{80, 90, 95, 78}
	fmt.Println(len(scores))
	fmt.Println(scores)

}
