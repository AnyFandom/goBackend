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

func (c Comments) Add(content string, postId uint) revel.Result {
	if !c.authorized {
		return c.RenderJsend("error", nil, "Not authorized")
	}
	comment := models.Comment{Content: content, PostID: postId, UserID: c.userId}
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
