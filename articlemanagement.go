package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Bio      string `json:"Bio"`
	Position string `json:"Position"`
}

var Articles []Article

func homePage(k http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(k, "Home!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(k http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(k).Encode(Articles)
}

func returnArticle(k http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(k).Encode(article)
		}
	}
}

func createArticle(k http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(k).Encode(article)
}

func deleteArticle(k http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	api := mux.NewRouter().StrictSlash(true)
	api.HandleFunc("/", homePage)
	api.HandleFunc("/view", returnAllArticles)
	api.HandleFunc("/create", createArticle).Methods("POST")
	api.HandleFunc("/delete/{id}", deleteArticle).Methods("DELETE")
	api.HandleFunc("/view/{id}", returnArticle)
	log.Fatal(http.ListenAndServe(":8888", api))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Name: "Kathir", Bio: "Trainee", Position: "Golang"},
		Article{Id: "2", Name: "Guru", Bio: "Unknown", Position: "Nothing"},
		Article{Id: "3", Name: "Scooby", Bio: "Dog", Position: "Pet1"},
		Article{Id: "4", Name: "Happy", Bio: "Dog", Position: "Pet2"},
		Article{Id: "5", Name: "Kutty", Bio: "Dog", Position: "Pet3"},
	}
	handleRequests()
}
