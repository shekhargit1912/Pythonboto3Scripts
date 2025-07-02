package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Grades []int
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func main() {
	student := Student{
		Name:   "Alice",
		Age:    20,
		Grades: []int{85, 90, 78},
	}

	fmt.Println("Student Name:", student.Name)
	fmt.Println("Student Age:", student.Age)
	fmt.Println("Average Grade:", student.AverageGrade())

	// Adding a new grade
	student.AddGrade(95)
	fmt.Println("Updated Grades:", student.Grades)
	fmt.Println("New Average Grade:", student.AverageGrade())
}
