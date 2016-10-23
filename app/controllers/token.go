package controllers

import (
	"goBackend/app/models"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Token struct {
	BaseController
}

func (c Token) Create(username string, password string) revel.Result {
	var user models.User
	c.Db.Where(&models.User{Name: username, Password: utils.HashPassword(password)}).First(&user)
	token := utils.CreateToken(user.ID)
	return c.RenderJsend("success", token)
}

func (c Token) Test() revel.Result {
	if !c.authorized {
		panic("You shall not pass!")
	}
	return c.RenderJsend("success", c.userId)
}
