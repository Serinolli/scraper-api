package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	Id      string
	Title   string
	Upvotes int
	Content string
}

func main() {
	resp, err := http.Get("http://localhost:8000/posts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Something went wrong!", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response []Post
	json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Error recovering data")
		return
	}

	fmt.Println(response)

}
