package controllers

import (
	"goBackend/app/models"

	"github.com/revel/revel"
)

type Comments struct {
	BaseController
}

func (c Comments) List() revel.Result {
	var comments []models.Comment
	c.Db.Find(&comments)
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

	comment := models.Comment{Content: content, PostID: postId, UserID: c.userId, ParentID: parent.ID, Depth: parent.Depth + 1}
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
	return c.RenderJsend("success", comment, "")
}
