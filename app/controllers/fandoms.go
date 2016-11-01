package controllers

import (
	"goBackend/app/models"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Fandoms struct {
	BaseController
}

func (c Fandoms) List() revel.Result {
	var fandoms []models.Fandom
	c.Db.Find(&fandoms)
	for _, v := range fandoms {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}
	return c.RenderJsend("success", fandoms, "")
}

func (c Fandoms) Add(title string, description string, avatar string) revel.Result {
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

	fandom := models.Fandom{Title: title, Description: description, Avatar: avatar}
	c.Db.NewRecord(fandom)
	c.Db.Create(&fandom)
	return c.RenderJsend("success", fandom, "")
}

func (c Fandoms) Item(id uint) revel.Result {
	var fandom models.Fandom
	c.Db.First(&fandom, id)
	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	return c.RenderJsend("success", fandom, "")
}

func (c Fandoms) ItemPosts(id uint) revel.Result {
	var fandom models.Fandom
	var posts []models.Post

	c.Db.First(&fandom, id)

	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	c.Db.Raw("SELECT * FROM posts WHERE blog_id in (SELECT id FROM blogs WHERE fandom_id = ?);", fandom.ID).Scan(&posts)

	for _, v := range posts {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}

	return c.RenderJsend("success", posts, "")
}

func (c Fandoms) ItemBlogs(id uint) revel.Result {
	var fandom models.Fandom
	var blogs []models.Blog

	c.Db.First(&fandom, id)

	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	c.Db.Where("fandom_id = ?", fandom.ID).Find(&blogs)

	for _, v := range blogs {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}

	return c.RenderJsend("success", blogs, "")
}

func (c Fandoms) ItemUpdate(id uint, title string, description string, avatar string) revel.Result {
	var fandom models.Fandom
	c.Db.First(&fandom, id)
	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Fandom not found")
	}
	if len(title) > 0 {
		c.Validation.Required(title)
		c.Validation.MaxSize(title, 100)
		c.Validation.MinSize(title, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		fandom.Title = title
	}
	if len(description) > 0 {
		c.Validation.Required(description)
		c.Validation.MaxSize(description, 100)
		c.Validation.MinSize(description, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		fandom.Description = description
	}
	if len(avatar) > 0 {
		c.Validation.Required(avatar)
		c.Validation.MaxSize(avatar, 100)
		c.Validation.MinSize(avatar, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		fandom.Avatar = avatar
	}
	c.Db.Save(&fandom)
	return c.RenderJsend("success", nil, "")
}

// TODO: Проверка авторизации
func (c Fandoms) ItemDelete(id uint) revel.Result {
	var fandom models.Fandom
	c.Db.First(&fandom, id)
	if fandom.ID == 0 {
		return c.RenderJsend("fail", nil, "Fandom not found")
	}
	c.Db.Debug().Delete(&fandom)
	return c.RenderJsend("success", nil, "")
}
