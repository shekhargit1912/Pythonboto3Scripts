package main

import "fmt"

func main() {
	var grades [5]int = [5]int{85, 90, 78, 92, 88}
	fmt.Println("Grades:", grades)
	fmt.Println("Length of grades array:", len(grades))

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i], "at index", i)
	}

	var fruites [5]string = [5]string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Println("Fruits:", fruites)

	games := [3]string{"Chess", "Football", "Tennis"}
	fmt.Println("Games:", games[2])

	for i, game := range games {
		fmt.Printf("Game %d: %s\n", i+1, game)
	}

	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length of numbers array:", len(numbers))

	array2D := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}

	fmt.Println("2D Array:")
	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D[i]); j++ {
			fmt.Println("Element at [", i, "][", j, "] is:", array2D[i][j])
		}
	}

}
