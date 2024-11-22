package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("printed")
	res, remainder, err := intDivision(5, 2)
	if err != nil {
		fmt.Print(err.Error())
	}

	switch remainder {
	case 0:
		fmt.Printf("Result: %v\n", res)
	default:
		fmt.Printf(`Result: %v
Remainder: %v`, res, remainder)
	}
}

func printMe(word string) {
	fmt.Println(word)
}

func intDivision(num int, denom int) (int, int, error) {
	var err error
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	res := num / denom
	remainder := num % denom
	return res, remainder, err
}
