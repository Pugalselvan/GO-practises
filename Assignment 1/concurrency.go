package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("jill")
}

func main() {
	go hello()
	time.Sleep(10 * time.Second)
	fmt.Println("Pugal")
}
