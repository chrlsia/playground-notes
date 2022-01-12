package main

import (
	"fmt"
	"myapp/house"
	"myapp/yard"
)

var i int
var num1 int = 10
var num2 int = 20
var flag bool = true

func main(){
	for i=1;i<=10;i++{
		switch i {
		case 2,4,6,8,10:
			printMessage(i)
			fmt.Println("Even")
		default:
			printMessage(i)
			fmt.Println("Odd")
		}
	}
	fmt.Print("Sum is ...")
	fmt.Println(addTwoIntegers(num1,num2))
	fmt.Print("Perimeter is ...")
	fmt.Println(house.Perimeter(1,2,3,4))

	house.Drawers()
	yard.HasLights(flag)

}

func printMessage(i int){
	fmt.Print(i," ")
}

/*
Notes:
command line: go main.go deck.go
(whatever relates to main only)
(function addTwoIntegers of deck.go is not exported so small letter)

myapp/house
(it meand import everything belonging to package house)
myapp/yard
(it means import everything belonging to package yard)

THAT's why their functions should be exported(=Capital first letter)
in order to be invoded from the main package
*/