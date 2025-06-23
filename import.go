package main

import (
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Harun Al Rasyid")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye)
}
