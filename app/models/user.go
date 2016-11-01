package models

import (
	"goBackend/app/utils"

	"github.com/jinzhu/gorm"
)

type User struct { // example user fields
	Username string `json:"username"`
	Password []byte `json:"-"`
	Base
}

func (m User) LoadInclude(Db *gorm.DB) []utils.IncludeItem {
	items := []utils.IncludeItem{}
	return items
}
