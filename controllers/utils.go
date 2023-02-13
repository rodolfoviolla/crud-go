package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodolfoviolla/crud-go/initializers"
	"github.com/rodolfoviolla/crud-go/models"
)

type Body struct {
	Body string
	Title string
}

func getBody(c *gin.Context) (*Body) {
	var body Body
	c.Bind(&body)
	return &body
}


func getPostByPostIdParam(c *gin.Context) (*models.Post) {
	var post models.Post
	postId := c.Param("postId")

	if result := initializers.DB.First(&post, postId); result.Error != nil {
		c.Status(400)
		return nil
	}

	return &post
}
