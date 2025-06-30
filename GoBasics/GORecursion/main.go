package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1 // Base case: factorial of 0 is 1
	}
	return n * factorial(n-1) // Recursive case
}

func namesss(ss string) {
	fmt.Println("Hello, ", ss)
}

func scores(score int) {
	fmt.Println("Your score is:", score)
}

// Anonymos function to calculate factorial

func main() {
	var number int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
	} else {
		result := factorial(number)
		fmt.Printf("Factorial of %d is %d\n", number, result)
	}

	fmt.Println("Anonymous function to calculate factorial:")
	factorialFunc := func(n int) int {
		if n == 0 {
			return 1 // Base case: factorial of 0 is 1
		}
		return n * n // Recursive case
	}
	resultAnon := factorialFunc(number)
	fmt.Printf("Factorial of %d using anonymous function is %d\n", number, resultAnon)
	fmt.Println("Recursion is a powerful technique in Go!")
	defer namesss("Alice")
	scores(95)
	fmt.Println("This is a simple example of recursion in Go.")

}
