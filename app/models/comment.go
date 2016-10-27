package models

type Comment struct {
	Content string
	UserID  uint
	PostID  uint
	Base
}
