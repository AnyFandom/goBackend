package models

import (
	"goBackend/app/utils"

	"github.com/jinzhu/gorm"
)

type Fandom struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Base
}

func (m Fandom) LoadInclude(Db *gorm.DB) []utils.IncludeItem {
	items := []utils.IncludeItem{}
	return items
}
