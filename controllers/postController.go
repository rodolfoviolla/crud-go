package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodolfoviolla/crud-go/initializers"
	"github.com/rodolfoviolla/crud-go/models"
)

func PostsCreate (c *gin.Context) {
	body := getBody(c)
	post := models.Post{Title: body.Title, Body: body.Body}
	
	if result := initializers.DB.Create(&post); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	if result := initializers.DB.Find(&posts); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}


func PostsShow(c *gin.Context) {
	post := getPostByPostIdParam(c)
	if post == nil { return }

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	post := getPostByPostIdParam(c)
	if post == nil { return }

	body := getBody(c)

	if result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	}); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	post := getPostByPostIdParam(c)
	if post == nil { return }

	if result := initializers.DB.Delete(&post); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{})
}