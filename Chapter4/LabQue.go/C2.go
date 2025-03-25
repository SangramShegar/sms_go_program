package main

import (
	"fmt"
	"sort"
)

type Student struct {
	rollno     int
	name       string
	percentage float64
}

type StudentList []Student

// Method to display students
func (s StudentList) Display() {
	for _, student := range s {
		fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f\n", student.rollno, student.name, student.percentage)
	}
}

// Sorting by percentage in descending order
func (s StudentList) SortByPercentage() {
	sort.Slice(s, func(i, j int) bool {
		return s[i].percentage > s[j].percentage
	})
}

func main() {
	students := StudentList{
		{101, "piyusha", 85.5},
		{102, "sakshi", 92.3},
		{103, "varsha", 89.4},
	}

	students.SortByPercentage()
	students.Display()
}
