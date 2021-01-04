package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcomw Pugal")

}

func handlRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8002", nil))
}
func main() {
	handlRequests()
}
