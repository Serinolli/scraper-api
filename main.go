package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string = "8000"

type Post struct {
	Id      string
	Title   string
	Upvotes int
	Content string
}

func main() {
	fmt.Println("Listening to port " + port)
	http.HandleFunc("/posts", getPosts)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPosts(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode([]Post{{
		Id:      "sampletest1",
		Title:   "starting api",
		Upvotes: 15,
		Content: "lorem ipsum dolor sit amet",
	}})
}
