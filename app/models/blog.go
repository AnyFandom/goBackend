package models

type Blog struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FandomID    uint   `json:"fandom_id"`
	Avatar      string `json:"avatar"`
	Base
}
