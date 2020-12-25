package main

import (
	"fmt"
)

const s string = "welcome"

func main() {
	fmt.Println(s)

	const n = 500000
	const d = 3e20 / n
	fmt.Println(d)
}
