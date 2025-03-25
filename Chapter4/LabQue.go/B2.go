package main

import "fmt"

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	case bool:
		fmt.Println("It's a bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	typeSwitch(42)
	typeSwitch("Hello")
	typeSwitch(true)
}
