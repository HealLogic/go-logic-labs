package main

import "fmt"

func main() {
	//one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 1

	//another way, declare type and name and assign value
	var secondNumber = 5

	//one step variable :  declare name, assign value and let go figure out the type
	otherNumber := 10

	fmt.Println("1 :", firstNumber)
	fmt.Println("2 :", secondNumber)
	fmt.Println("3 :", otherNumber)
}