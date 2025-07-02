// Learming Go Interfaces
// This code demonstrates the use of interfaces in Go.

package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	ract := Rectangle{Width: 5, Height: 10}
	printShapeInfo(ract)

	fmt.Println("Rectangle Area:", ract.Area())
	fmt.Println("Rectangle Perimeter:", ract.Perimeter())
	fmt.Println("Shape interface implemented successfully.")
	fmt.Println("This is an example of using interfaces in Go.")
	fmt.Println("Interfaces allow us to define a contract for types to implement.")
	fmt.Println("We can use interfaces to write more generic and reusable code.")

	fmt.Println("In this example, we defined a Shape interface and implemented it for the Rectangle type.")
}
