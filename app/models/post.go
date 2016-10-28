package models

type Post struct {
	Title   string
	Content string
	UserID  uint
	BlogID  uint
	Base
}
