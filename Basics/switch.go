package main

import "fmt"

func main() {
	age := 22

	switch age {
	case 16:
		fmt.Println("went to ooty")
	case 18:
		fmt.Println("went to pondy")
	case 22:
		fmt.Println("went to wayanad")
	default:
		fmt.Println("im in home")

	}
}
