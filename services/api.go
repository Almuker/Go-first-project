package services

import (
	"Go_project/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := models.Articles {
		models.Article{Title: "Test title", Desc: "Test desc", Content: "Test content"},
	}

	fmt.Println("Endpoint Hit: All Articles Enpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePageHandler)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	fmt.Println("Server listening on port 8080")
	log.Panic(http.ListenAndServe(":8080", myRouter))
}