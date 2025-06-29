package main

import "fmt"

func main() {
	slices := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slices)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[3:7] // Slice from index 2 to 4
	fmt.Println("Slice s1:", s1)

	srs := make([]int, 5, 8) // Create a slice with length 5
	fmt.Println(srs)
	fmt.Println(len(srs))
	fmt.Println(cap(srs)) // Capacity is 8

	// Append elements to the slice

	arr1 := []int{1, 2, 3, 5, 6, 7, 8, 9}
	slice1 := arr1[1:3]
	fmt.Println("Slice from arr1:", slice1)
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	slice1 = append(slice1, 10, 11, 12) // Append elements
	fmt.Println("Slice after appending:", slice1)
	fmt.Println("Length after appending:", len(slice1))
	fmt.Println("Capacity after appending:", cap(slice1))

}
