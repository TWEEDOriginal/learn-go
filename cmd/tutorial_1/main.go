package main

import (
	"fmt"
	"math"
)

func main() {
	var bread string
	fmt.Print("your word: ")
	fmt.Scan(&bread)
	fmt.Println(math.Pow(3, 2), 5, bread)
}
