package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Println("Address of a variable: %x\n", &a)
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}
