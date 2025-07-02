package main

import "fmt"

func modifyValue(x *int) {
	*x = 20 // Modifying the value at the address pointed to by x
}

func modigyone(x int) {
	fmt.Println("Value of x in modifyone:", x)
	x = 30 // This will not affect the original variable
	fmt.Println("Modified value of x in modifyone:", x)
}

func main() {

	var a int = 10
	fmt.Println("Original value of a:", a)
	modifyValue(&a) // Passing the address of a to the function
	fmt.Println("Modified value of a:", a)

	modigyone(a)                                  // Passing the value of a to the function
	fmt.Println("Value of a after modigyone:", a) // a remains unchanged
	fmt.Println("This is an example of passing by value and passing by reference in Go.")

}
