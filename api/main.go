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
