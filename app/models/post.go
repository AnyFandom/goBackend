package models

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id"`
	Base
}
