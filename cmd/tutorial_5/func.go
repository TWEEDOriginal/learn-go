package main

import "fmt"

type owner struct {
	name string
}

type gasEngine struct {
	mpg       uint16
	gallons   uint16
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint16
	kwh   uint16
}

// method for gasEngine struct
func (e gasEngine) milesLeft() uint16 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint16 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint16
}

func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Println("Can reach destination")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"tweed"}}
	fmt.Println(myEngine)
	fmt.Printf("Total miles left in gasoline tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
	var mySecondEngine electricEngine = electricEngine{25, 15}
	canMakeIt(mySecondEngine, 500)
}
