package main

import "fmt"

func main() {
	age := 32
	// use ampersand to get memory address
	agePointer := &age
	// use asterisk if you want to get the actual value
	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
}

// input is memory address
func editAgeToAdultYears(age *int) {
	// overwrite value stored in age
	*age = *age - 18
}
