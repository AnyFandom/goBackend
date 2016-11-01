package controllers

import (
	"goBackend/app/models"
	"goBackend/app/utils"

	"github.com/revel/revel"
)

type Comments struct {
	BaseController
}

func (c Comments) List() revel.Result {
	var comments []models.Comment
	c.Db.Find(&comments)
	for _, v := range comments {
		c.include = utils.ExtendInclude(c.include, v.LoadInclude(c.Db))
	}
	return c.RenderJsend("success", comments, "")
}

func (c Comments) Add(content string, postId uint, parentId uint) revel.Result {
	if !c.authorized {
		return c.RenderJsend("error", nil, "Not authorized")
	}

	c.Validation.Required(content)
	c.Validation.MaxSize(content, 255)
	c.Validation.MinSize(content, 2)

	if c.Validation.HasErrors() {
		return c.RenderJsend("fail", nil, "Validation error")
	}

	var post models.Post
	c.Db.First(&post, postId)
	if post.ID == 0 {
		return c.RenderJsend("fail", nil, "Post not found")
	}

	var parent models.Comment
	c.Db.First(&parent, parentId)

	if parent.ID != 0 && parent.PostID != post.ID {
		return c.RenderJsend("fail", nil, "Parent in other post")
	}

	var depth int

	if parent.ID == 0 {
		depth = 0
	} else {
		depth = parent.Depth + 1
	}

	comment := models.Comment{Content: content, PostID: postId, UserID: c.userId, ParentID: parent.ID, Depth: depth}
	c.Db.NewRecord(comment)
	c.Db.Create(&comment)
	return c.RenderJsend("success", comment, "")
}

func (c Comments) Item(id uint) revel.Result {
	var comment models.Comment
	c.Db.First(&comment, id)
	if comment.ID == 0 {
		return c.RenderJsend("fail", nil, "Comment not found")
	}

	c.include = utils.ExtendInclude(c.include, comment.LoadInclude(c.Db))

	return c.RenderJsend("success", comment, "")
}

func (c Comments) ItemUpdate(id uint, content string) revel.Result {
	var comment models.Comment
	c.Db.First(&comment, id)
	if comment.ID == 0 {
		return c.RenderJsend("fail", nil, "Comment not found")
	}
	if len(content) > 0 {
		c.Validation.Required(content)
		c.Validation.MaxSize(content, 255)
		c.Validation.MinSize(content, 2)

		if c.Validation.HasErrors() {
			return c.RenderJsend("fail", nil, "Validation error")
		}

		comment.Content = content
	}
	c.Db.Save(&comment)
	return c.RenderJsend("success", nil, "")
}

// TODO: Проверка авторизации
func (c Comments) ItemDelete(id uint) revel.Result {
	var comment models.Comment
	c.Db.First(&comment, id)
	if comment.ID == 0 {
		return c.RenderJsend("fail", nil, "Comment not found")
	}
	c.Db.Debug().Delete(&comment)
	return c.RenderJsend("success", nil, "")
}
