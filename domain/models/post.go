package model

import "time"

type Post struct {
	PostId    string    `json:"postid"`
	SubReddit string    `json:"subReddit"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Author    string    `json:"author"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Upvotes   int       `json:"upvotes"`
}
