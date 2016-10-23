package controllers

import (
	"goBackend/app/models"
	"goBackend/app/routes"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Users struct {
	BaseController
}

func (c Users) List() revel.Result {
	var users []models.User
	c.Db.Find(&users)
	return c.RenderJsend("success", users)
}

func (c Users) Item(id uint) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	return c.RenderJsend("success", user)
}

func (c Users) Add(username string, password string) revel.Result {
	var user = models.User{Name: username, Password: utils.HashPassword(password)}
	c.Db.NewRecord(user)
	c.Db.Create(&user)
	var location = utils.Location{Location: routes.Users.Item(user.ID)}
	return c.RenderJsend("success", location)
}
