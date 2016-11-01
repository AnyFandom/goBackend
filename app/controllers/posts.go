package controllers

import (
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
	for _, v := range posts {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}
	return c.RenderJsend("success", posts, "")
}

func (c Posts) Item(id uint) revel.Result {
	var post models.Post
	c.Db.First(&post, id)

	c.include = utils.ExtendInclude(c.include, post.LoadInclude(c.Db))

	return c.RenderJsend("success", post, "")
}

func (c Posts) Add(title string, content string, blogId uint) revel.Result {
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

	var blog models.Blog
	c.Db.First(&blog, blogId)
	if blog.ID == 0 {
		return c.RenderJsend("fail", nil, "Blog not found")
	}

	var user models.User
	c.Db.First(&user, c.userId)

	var post = models.Post{Title: title, Content: content, UserID: user.ID, BlogID: blog.ID}

	c.Db.NewRecord(post)
	c.Db.Create(&post)

	var location = utils.Location{Location: routes.Posts.Item(post.ID)}

	return c.RenderJsend("success", location, "")
}

func (c Posts) ItemComments(id uint) revel.Result {
	var post models.Post
	c.Db.First(&post, id)
	if post.ID == 0 {
		return c.RenderJsend("fail", nil, "Not found")
	}

	var comments []models.Comment
	c.Db.Where(&models.Comment{PostID: id}).Find(&comments)
	for _, v := range comments {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}
	return c.RenderJsend("success", comments, "")
}

func (c Posts) ItemUpdate(id uint, title string, content string, avatar string) revel.Result {
	var post models.Post
	c.Db.First(&post, id)
	if post.ID == 0 {
		return c.RenderJsend("fail", nil, "Post not found")
	}
	if len(title) > 0 {
		c.Validation.Required(title)
		c.Validation.MaxSize(title, 100)
		c.Validation.MinSize(title, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		post.Title = title
	}
	if len(content) > 0 {
		c.Validation.Required(content)
		c.Validation.MaxSize(content, 100)
		c.Validation.MinSize(content, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		post.Content = content
	}
	c.Db.Save(&post)
	return c.RenderJsend("success", nil, "")
}

// TODO: Проверка авторизации
func (c Posts) ItemDelete(id uint) revel.Result {
	var post models.Post
	c.Db.First(&post, id)
	if post.ID == 0 {
		return c.RenderJsend("fail", nil, "Post not found")
	}
	c.Db.Debug().Delete(&post)
	return c.RenderJsend("success", nil, "")
}
