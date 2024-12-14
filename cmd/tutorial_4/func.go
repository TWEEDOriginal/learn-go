package main

import (
	"fmt"
	"strings"
)

func main() {
	// len is 8 because é uses 2 bytes
	// uses variable length encoding
	// number of prefix 1s of first byte == no of bytes symbol uses
	var mystring1 = "résumé"
	// returns only first half of é
	var indexed = mystring1[1]
	// value (int) and type
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range mystring1 {
		fmt.Println(i, v)
	}
	fmt.Printf("\nLen of old string is %v", len(mystring1))
	// convert to rune array for proper indexing with len of chars
	var mystring2 = []rune(mystring1)
	fmt.Printf("\nLen of new string is %v", len(mystring2))

	var strSlice = []string{"s", "u", "p", "e", "r"}
	// faster for creating strings
	// arr alocated internally
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	// string is then formed from arr
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
