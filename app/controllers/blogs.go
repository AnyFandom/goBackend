package controllers

import (
	"goBackend/app/models"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Blogs struct {
	BaseController
}

func (c Blogs) List() revel.Result {
	var blogs []models.Blog
	c.Db.Find(&blogs)
	for _, v := range blogs {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}
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

	c.include = utils.ExtendInclude(c.include, blog.LoadInclude(c.Db))

	c.ExtendInclude(blog.LoadInclude(c.Db))
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

	for _, v := range posts {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}

	return c.RenderJsend("success", posts, "")
}

func (c Blogs) ItemUpdate(id uint, title string, description string, avatar string) revel.Result {
	var blog models.Blog
	c.Db.First(&blog, id)
	if blog.ID == 0 {
		return c.RenderJsend("fail", nil, "Blog not found")
	}
	if len(title) > 0 {
		c.Validation.Required(title)
		c.Validation.MaxSize(title, 100)
		c.Validation.MinSize(title, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		blog.Title = title
	}
	if len(description) > 0 {
		c.Validation.Required(description)
		c.Validation.MaxSize(description, 100)
		c.Validation.MinSize(description, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		blog.Description = description
	}
	if len(avatar) > 0 {
		c.Validation.Required(avatar)
		c.Validation.MaxSize(avatar, 100)
		c.Validation.MinSize(avatar, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		blog.Avatar = avatar
	}
	c.Db.Save(&blog)
	return c.RenderJsend("success", nil, "")
}

// TODO: Проверка авторизации
func (c Blogs) ItemDelete(id uint) revel.Result {
	var blog models.Blog
	c.Db.First(&blog, id)
	if blog.ID == 0 {
		return c.RenderJsend("fail", nil, "Blog not found")
	}
	c.Db.Debug().Delete(&blog)
	return c.RenderJsend("success", nil, "")
}
