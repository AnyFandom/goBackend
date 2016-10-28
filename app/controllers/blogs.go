package controllers

import (
	"goBackend/app/models"

	"github.com/revel/revel"
)

type Blogs struct {
	BaseController
}

func (c Blogs) List() revel.Result {
	var blogs []models.Blog
	c.Db.Find(&blogs)
	return c.RenderJsend("success", blogs, "")
}

func (c Blogs) Add(title string, description string, avatar string, fandomId uint) revel.Result {
	if !c.authorized {
		return c.RenderJsend("fail", nil, "Not authorized")
	}

	c.Validation.Required(title)
	c.Validation.MaxSize(title, 100)
	c.Validation.MinSize(title, 2)

	c.Validation.Required(description)
	c.Validation.MaxSize(description, 100)
	c.Validation.MinSize(description, 2)

	if c.Validation.HasErrors() {
		return c.RenderJsend("fail", nil, "Validation error")
	}

	var fandom models.Fandom
	c.Db.First(&fandom, fandomId)
	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Fandom not found")
	}

	blog := models.Blog{Title: title, Description: description, Avatar: avatar, FandomID: fandom.ID}
	c.Db.NewRecord(blog)
	c.Db.Create(&blog)
	return c.RenderJsend("success", blog, "")
}

func (c Blogs) Item(id uint) revel.Result {
	var blog models.Blog
	c.Db.First(&blog, id)
	if blog.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	return c.RenderJsend("success", blog, "")
}

func (c Blogs) ItemPosts(id uint) revel.Result {
	var blog models.Blog
	var posts []models.Post

	c.Db.First(&blog, id)

	if blog.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	c.Db.Where("blog_id = ?", blog.ID).Find(&posts)

	return c.RenderJsend("success", posts, "")
}
