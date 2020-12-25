package main

import "fmt"

func main() {
	strName := map[string]string{"jill": "boosh", "vaishy": "bubo"}
	for index, element := range strName {
		fmt.Println("Index :", index, " Element :", element)
	}

	for key := range strName {
		fmt.Println(key)
	}
}
