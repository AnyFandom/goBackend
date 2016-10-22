package controllers

import (
	"goBackend/app/models"

	"github.com/revel/revel"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
	user := models.User{Name: "Jinzhup"}
	c.Txn.NewRecord(user)
	c.Txn.Create(&user)
	return c.RenderJson(user)
}
