package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type articles []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	allArticles := articles{
		article{Title: "master", Description: "vijay"},
		article{Title: "valimai", Description: "ajith"},
		article{Title: "vtv", Description: "simbu"},
		article{Title: "run", Description: "vikram"},
		article{Title: "maara", Description: "soorya"},
	}

	fmt.Println("end points hit all the articles")
	json.NewEncoder(w).Encode(allArticles)
}

func homHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page endpoint hits")
}

func requestHandler() {
	http.HandleFunc("/", homHandler)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8004", nil))
}

func main() {
	requestHandler()
}
