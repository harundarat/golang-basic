package main

import "fmt"

func main() {
	type NoKTP string

	const harun NoKTP = "1000"
	fmt.Println(harun)

	const rasyid string = "2000"
	const KTPRasyid NoKTP = NoKTP(rasyid)
	fmt.Println(KTPRasyid)
}
