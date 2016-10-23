package controllers

import (
	"goBackend/app/models"
	"goBackend/app/routes"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

// TODO: Add jsend

type Users struct {
	BaseController
}

func (c Users) List() revel.Result {
	var users []models.User
	c.Txn.Find(&users)
	return c.RenderJson(users)
}

func (c Users) Item(id int64) revel.Result {
	var user models.User
	c.Txn.First(&user, id)
	return c.RenderJsend("success", user)
}

func (c Users) Add(username string, password string) revel.Result {
	var user = models.User{Name: username, Password: utils.HashPassword(password)}
	c.Txn.NewRecord(user)
	c.Txn.Create(&user)
	var location = utils.Location{Location: routes.Users.Item(int64(user.ID))}
	return c.RenderJsend("success", location)
}
