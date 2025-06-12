package main

import "fmt"

func main() {
	var value32 int32 = 32000
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)
	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	var name = "Harun Al Rasyid"
	var e uint8 = name[1]
	var eString = string(e)

	fmt.Println(eString)

}
