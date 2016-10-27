package controllers

import (
	"fmt"
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
	return c.RenderJsend("success", users, "")
}

func (c Users) Item(id uint) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	return c.RenderJsend("success", user, "")
}

func (c Users) Add(username string, password string) revel.Result {
	c.Validation.Required(username)
	c.Validation.MaxSize(username, 15)
	c.Validation.MinSize(username, 2)

	c.Validation.Required(password)
	c.Validation.MaxSize(password, 200)
	c.Validation.MinSize(password, 6)

	if c.Validation.HasErrors() {
		return c.RenderJsend("fail", nil, "Validation error")
	}

	var user = models.User{Username: username, Password: utils.HashPassword(password)}
	c.Db.NewRecord(user)
	c.Db.Create(&user)
	var location = utils.Location{Location: routes.Users.Item(user.ID)}
	return c.RenderJsend("success", location, "")
}

func (c Users) Current() revel.Result {
	if !c.authorized {
		return c.RenderJsend("fail", nil, "Not authorized")
	}
	fmt.Println(c.authorized, c.userId)
	var user models.User
	c.Db.First(&user, c.userId)
	return c.RenderJsend("success", user, "")
}

func (c Users) ItemPosts(id uint) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}

	var posts []models.Post
	c.Db.Where(&models.Post{UserID: id}).Find(&posts)
	return c.RenderJsend("success", posts, "")
}

func (c Users) CurrentPosts() revel.Result {
	var user models.User
	c.Db.First(&user, c.userId)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}

	var posts []models.Post
	c.Db.Where(&models.Post{UserID: c.userId}).Find(&posts)
	return c.RenderJsend("success", posts, "")
}
