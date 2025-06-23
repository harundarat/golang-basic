package main

import (
	"errors"
	"fmt"
)

func multiplication(number int, multiplier int) (int, error) {
	if multiplier == 0 || number == 0 {
		return 0, errors.New("Multiplication with zero")
	} else {
		return number * multiplier, nil
	}
}

func main() {
	hasil, err := multiplication(5, 0)

	if err == nil {
		fmt.Print("Result: ", hasil)
	} else {
		fmt.Println("Error: ", err)
	}
}
