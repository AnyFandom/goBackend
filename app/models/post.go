package models

import (
	"goBackend/app/utils"

	"github.com/jinzhu/gorm"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id"`
	Base
}

func (m Post) LoadInclude(Db *gorm.DB) []utils.IncludeItem {
	items := []utils.IncludeItem{}

	var blog Blog
	Db.First(&blog, m.BlogID)
	items = append(items, utils.IncludeItem{Type: "blog", Data: blog})

	var user User
	Db.First(&user, m.UserID)
	items = append(items, utils.IncludeItem{Type: "user", Data: user})

	items = utils.ExtendInclude(items, blog.LoadInclude(Db))
	items = utils.ExtendInclude(items, user.LoadInclude(Db))

	return items
}
