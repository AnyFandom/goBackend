package controllers

import (
	"fmt"
	"goBackend/app/models"
	"goBackend/app/routes"
	"goBackend/app/utils"
	"reflect"

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
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}
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

func (c Users) ItemComments(id uint) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}

	var comments []models.Comment
	c.Db.Where(&models.Comment{UserID: id}).Find(&comments)
	return c.RenderJsend("success", comments, "")
}

func (c Users) CurrentComments() revel.Result {
	var user models.User
	c.Db.First(&user, c.userId)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}

	var comments []models.Comment
	c.Db.Where(&models.Comment{UserID: c.userId}).Find(&comments)
	return c.RenderJsend("success", comments, "")
}

func (c Users) ItemUpdate(id uint, username string, password string, password_old string) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}
	if len(username) > 0 {
		c.Validation.Required(username)
		c.Validation.MaxSize(username, 100)
		c.Validation.MinSize(username, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		user.Username = username
	}

	if len(password) > 0 {
		c.Validation.Required(password)
		c.Validation.MaxSize(password, 200)
		c.Validation.MinSize(password, 6)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		if len(password_old) == 0 {
			return c.RenderJsend("fail", nil, "No old password")
		}
		if !reflect.DeepEqual(utils.HashPassword(password_old), user.Password) {
			return c.RenderJsend("fail", nil, "Old password incorrect")
		}

		user.Password = utils.HashPassword(password)
	}
	c.Db.Save(&user)
	return c.RenderJsend("success", nil, "")
}

// TODO: Проверка авторизации
func (c Users) ItemDelete(id uint) revel.Result {
	var user models.User
	c.Db.First(&user, id)
	if user.ID == 0 {
		return c.RenderJsend("fail", nil, "User not found")
	}
	c.Db.Debug().Delete(&user)
	return c.RenderJsend("success", nil, "")
}
