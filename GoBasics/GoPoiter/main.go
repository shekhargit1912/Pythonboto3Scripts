package main

import "fmt"

func main() {
	// Variable declaration
	i := 11
	var ptr_i *int = &i // Pointer to i
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Pointer to i:", *ptr_i)

	s := "shekhar"
	var ptr_s = &s // Pointer to s
	fmt.Println("Value of s:", *ptr_s)

	s1 := "shekhar"
	var ptr_s1 *string = &s1 // Pointer to s1
	fmt.Println("Value of s1:", *ptr_s1)

	// short variable declaration
	fmt.Printf("Value of i: %T %v ", &i, &i)
	fmt.Printf("Value of i: %T %v", *(&i), *(&i))

}
