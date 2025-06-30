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

	//Xopy from the slice

	src_slice := []int{1, 2, 3, 4, 5}
	dst_slice := make([]int, 3)       // Create a destination slice with the same length
	num := copy(dst_slice, src_slice) // Copy elements from src_slice to dst_slice
	fmt.Println("Source slice:", src_slice)
	fmt.Println("Destination slice after copy:", dst_slice)
	fmt.Println("Number of elements copied:", num)

	//Map in go
	demo := map[string]int{"A": 65, "B": 66, "C": 67}
	delete(demo, "B") // Delete key "B" from the map
	fmt.Println("Demo map after deletion:", demo)

	/*
		- create a map using make() function with key data type as string, and value data type as int.
		- add the following key_value pairs to it
		("A", 65)
		("F", 70)
		("K", 75)
		- delete the key "F"
		- print the map
	*/

	// Creating a map using make() function
	// with key data type as string, and value data type as int.
	// Adding key-value pairs
	demoMap := make(map[string]int)
	demoMap["A"] = 65
	demoMap["F"] = 70
	demoMap["K"] = 75
	// Deleting the key "F"
	delete(demoMap, "F")
	// Printing the map
	fmt.Println("Demo map after adding and deleting:", demoMap)

	codes := map[string]int{
		"Go":     1,
		"Python": 2,
		"Java":   3}

	fmt.Println("Codes map:", codes)
	fmt.Println(codes["Go"])
	fmt.Println("Length of codes map:", len(codes))

	// gettinmg keys
	value, found := codes["Go"]
	if found {
		fmt.Println("Value for key 'Go':", found, value)
	} else {
		fmt.Println("Key 'Go' not found in map")
	}
	// getting value for a key
	value, found = codes["Javas"]
	if found {
		fmt.Println("Value for key 'Javas':", found, value)
	} else {
		fmt.Println("Key 'Javas' not found in map")
	}

	//adding a new key-value pair
	codes["JavaScript"] = 4
	fmt.Println("Updated codes map after adding JavaScript:", codes)
	// deleting a key-value pair
	delete(codes, "Python")
	fmt.Println("Codes map after deleting Python:", codes)

	// iterating over a map
	for key, value := range codes {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//truue map example
	// True map example
	// codes2 := map[string]int{"Go": 1, "Python": 2, "Java": 3}
	// fmt.Println("True map example:", codes2)
	// fmt.Println("Value for key 'Go':", codes2["Go"])
	// fmt.Println("Length of true map example:", len(codes2))
	// fmt.Println("Keys in true map example:")
	// for key := range codes2 {
	// 	fmt.Println(key)

	// }

	//fmt.Println("True map example:")

	// initiualize empty map
	codes1 := make(map[string]int)
	codes1["Go"] = 1
	fmt.Println("Empty map after adding Go:", codes1)

}
