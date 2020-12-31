package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	boo := User{"ram", "jill@boo.bu", 22}
	fmt.Printf("%v+\n", boo)
	fmt.Printf("%v+\n", boo.Email)

	var sat = new(User)
	sat.Name = "selva"
	sat.Email = "boo@loc.bu"
	sat.age = 23
	fmt.Printf("%v+\n", sat)
}
