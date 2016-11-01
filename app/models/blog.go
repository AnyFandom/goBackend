package models

import (
	"goBackend/app/utils"

	"github.com/jinzhu/gorm"
)

type Blog struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FandomID    uint   `json:"fandom_id"`
	Avatar      string `json:"avatar"`
	Base
}

func (m Blog) LoadInclude(Db *gorm.DB) []utils.IncludeItem {
	items := []utils.IncludeItem{}
	var fandom Fandom
	Db.First(&fandom, m.FandomID)
	items = append(items, utils.IncludeItem{Type: "fandom", Data: fandom})

	items = utils.ExtendInclude(items, fandom.LoadInclude(Db))

	return items
}
