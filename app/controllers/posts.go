package controllers

import (
	"fmt"
	"goBackend/app/models"
	"goBackend/app/routes"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Posts struct {
	BaseController
}

func (c Posts) List() revel.Result {
	var posts []models.Post
	c.Db.Find(&posts)
	return c.RenderJsend("success", posts, "")
}

func (c Posts) Item(id uint) revel.Result {
	var post models.Post
	c.Db.First(&post, id)
	return c.RenderJsend("success", post, "")
}

func (c Posts) Add(title string, content string) revel.Result {
	if !c.authorized {
		return c.RenderJsend("fail", nil, "Not authorized")
	}

	c.Validation.Required(title)
	c.Validation.MaxSize(title, 100)
	c.Validation.MinSize(title, 2)

	c.Validation.Required(content)
	c.Validation.MaxSize(content, 255)
	c.Validation.MinSize(content, 2)

	if c.Validation.HasErrors() {
		return c.RenderJsend("fail", nil, "Validation error")
	}

	var user models.User
	c.Db.First(&user, c.userId)

	var post = models.Post{Title: title, Content: content, UserID: user.ID}
	fmt.Println(post)
	c.Db.NewRecord(post)
	c.Db.Create(&post)

	var location = utils.Location{Location: routes.Posts.Item(post.ID)}

	return c.RenderJsend("success", location, "")
}
