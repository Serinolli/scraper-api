package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/Serinolli/scraper-api/models"
	"github.com/gorilla/mux"
)

var port string = "8000"

func main() {
	muxRouter := mux.NewRouter()

	fmt.Println("Listening to port " + port)
	muxRouter.HandleFunc("/posts", getPosts).Methods("GET")
	muxRouter.HandleFunc("/posts", createPost).Methods("POST")
	//muxRouter.HandleFunc("/posts/{id}", getPost).Methods("GET")
	//muxRouter.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	//muxRouter.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, muxRouter))
}

func getPosts(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode([]models.Post{{
		ID:      "sampletest1",
		Title:   "starting api",
		Upvotes: 15,
		Content: "lorem ipsum dolor sit amet",
	}})
}

func createPost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var posts []models.Post
	err := json.NewDecoder(request.Body).Decode(&posts)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(writer).Encode(posts)
}
