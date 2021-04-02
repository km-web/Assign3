package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(k http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(k, "Kathir Batch11!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	api := mux.NewRouter().StrictSlash(true)
	api.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8888", api))
}

func main() {

	handleRequests()
}
