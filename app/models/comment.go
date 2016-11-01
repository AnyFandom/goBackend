package models

import (
	"goBackend/app/utils"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	PostID   uint   `json:"post_id"`
	Depth    int    `json:"depth"`
	ParentID uint   `json:"parent_id"`
	Base
}

func (m Comment) LoadInclude(Db *gorm.DB) []utils.IncludeItem {
	items := []utils.IncludeItem{}

	var post Post
	Db.First(&post, m.PostID)
	items = append(items, utils.IncludeItem{Type: "post", Data: post})

	var user User
	Db.First(&user, m.UserID)
	items = append(items, utils.IncludeItem{Type: "user", Data: user})

	items = utils.ExtendInclude(items, post.LoadInclude(Db))
	items = utils.ExtendInclude(items, user.LoadInclude(Db))

	return items
}
