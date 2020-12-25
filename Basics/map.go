package main

import "fmt"

func main() {
	employee := make(map[string]int)

	employee["jill"] = 22
	employee["boosh"] = 23
	employee["raw"] = 24
	employee["robin"] = 30

	fmt.Println(len(employee))
}
