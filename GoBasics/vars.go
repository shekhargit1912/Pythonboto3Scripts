package main

import (
	"fmt"
	"strconv"
)

func vars() {

	var i int = 10
	var s string = strconv.Itoa(i)
	fmt.Printf("Integer i as string: %q \n", s)

	var ss string = "100abv"
	ii, err := strconv.Atoi(ss)
	fmt.Printf("String ss as integer: %v, error: %T \n", ii, ii)
	fmt.Printf("String ss as integer: %v, error: %T \n", err, err)
	fmt.Printf("String ss as integer: %d, error: %v \n", ii, err)

	citty := "Istanbul"
	{
		// Variable declaration
		var city string = "Ankara"
		println("City inside block:", city)
		println("City outside block:", citty)
	}
	fmt.Println("City outside block:", citty)
	fmt.Println("Variable city is not accessible outside the block")

	//user input
	var name string
	var AGE int
	var is_muggle bool
	fmt.Print("Enter your name: AND AGE ")
	// fmt.Scanf("%s %d %t", &name, &AGE, &is_muggle)
	fmt.Println("Hello,", name, "!")
	fmt.Println("Your age is:", AGE)
	fmt.Println("Are you a muggle?", is_muggle)

}

func main() {
	vars()
}
