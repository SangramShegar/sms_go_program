package main

import "fmt"

type Student struct {
	name  string
	roll  int
	grade string
}
func (s *Student) Show() {
	fmt.Printf("Name: %s, Roll No: %d, Grade: %s\n", s.name, s.roll, s.grade)
}

func main() {
	student := &Student{name: "piyusha", roll: 106, grade: "A"}
	student.Show()
}
