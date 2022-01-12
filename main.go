package main

import "fmt"

var i int
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
}

func printMessage(i int){
	fmt.Print(i," ")
}