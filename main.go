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