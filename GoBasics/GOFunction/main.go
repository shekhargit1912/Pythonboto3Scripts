package main

import "fmt"

func addnum(a int, b int) int {

	sum := a + b
	return sum
}

func operation(s int, t int) (sum1 int, diff int) {
	// This function returns two values: sum and difference
	sum1 = s + t
	diff = s - t
	return // Implicit return of sum1 and diff
}

func noReturnValue(studentName string, subjects ...string) {
	// This function takes a string and a variable number of subjects
	fmt.Println("Subjects for", studentName, "are:")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	// This function does not return any value
	fmt.Println("Student Name:", studentName)
}

func main() {
	var a int = 10
	var b int = 20
	sum := addnum(a, b)
	fmt.Println("Sum of", a, "and", b, "is:", sum)

	sum1, diff := operation(30, 10)
	fmt.Println("Sum:", sum1, "Difference:", diff)

	noReturnValue("Alice", "Math", "Science", "History")

	// Example of a function with no return value
}
