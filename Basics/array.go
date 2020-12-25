package main

import "fmt"

func main() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 200
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("values[%d] = %d\n", j, n[j])
	}
}
