package main

import "fmt"

func main() {
	names := [...]string{"Harun", "Al", "Rasyid", "We", "Love", "Golang", "<3"}

	slice1 := names[3:6]
	slice2 := names[3:]
	slice3 := names[:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "new saturday"
	daySlice1[1] = "new sunday"
	fmt.Println("days: ", days)
	fmt.Println("daySlice1: ", daySlice1)

	daySlice2 := append(daySlice1, "new day") // this will create new array inside new slice, since length of array `days` can't be changed
	fmt.Println("days: ", days)
	fmt.Println("daySlice1: ", daySlice1)
	fmt.Println("daySlice2: ", daySlice2)

	daySlice2[0] = "old saturday" // this won't affect `days` because the array already different
	fmt.Println("days: ", days)
	fmt.Println("daySlice1: ", daySlice1)
	fmt.Println("daySlice2:", daySlice2)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "harun"
	newSlice[1] = "harun"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// newSlice[2] = "harun" // this won't work because out of range. Must use append
	newSlice2 := append(newSlice, "harun")
	fmt.Println("newSlice2: ", newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "rasyid" // this will affect `newSlice` because still using original array
	fmt.Println("newSlice2: ", newSlice2)
	fmt.Println("newSlice: ", newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("fromSlice: ", fromSlice)
	fmt.Println("toSlice: ", toSlice)

	thisIsArray := [...]string{"Harun", "Al", "Rasyid"}
	thisIsSlice := []string{"Harun", "Al", "Rasyid"}
	fmt.Println("thisIsArray: ", thisIsArray)
	fmt.Println("thisIsSlice: ", thisIsSlice)
}
