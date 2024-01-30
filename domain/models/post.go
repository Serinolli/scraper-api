package model

type Post struct {
	PostId    string `json:"postid"`
	SubReddit string `json:"subReddit"`
	Timestamp int    `json:"timeStamp"`
	Author    string `json:"author"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Upvotes   int    `json:"upvotes"`
	Content   string `json:"content"`
}
