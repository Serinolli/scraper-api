package model

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Upvotes int    `json:"upvotes"`
	Content string `json:"content"`
}
