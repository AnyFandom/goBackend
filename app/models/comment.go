package models

type Comment struct {
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	PostID   uint   `json:"post_id"`
	Depth    int    `json:"depth"`
	ParentID uint   `json:"parent_id"`
	Base
}
