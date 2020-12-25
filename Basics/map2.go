package main

import "fmt"

func main() {
	superhero := map[string]map[string]string{
		"dev": map[string]string{
			"name": "jill",
			"city": "cbe",
		},
		"sales": map[string]string{
			"name": "boosh",
			"city": "banglore",
		},
	}
	if temp, hero := superhero["dev"]; hero {
		fmt.Println(temp["name"], temp["city"])
	}
}
