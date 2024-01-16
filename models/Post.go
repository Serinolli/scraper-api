package model

import "go.mongodb.org/mongo-driver/mongo"

type Post struct {
	ID        string `json:"id"`
	SubReddit string `json:"subReddit"`
	Timestamp string `json:"timeStamp"`
	Author    string `json:"author"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Upvotes   int    `json:"upvotes"`
	Content   string `json:"content"`
	Client    *mongo.Client
}

func NewPostWorker(c *mongo.Client) *Post {
	return &Post{
		Client: c,
	}
}
