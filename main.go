package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodolfoviolla/crud-go/controllers"
	"github.com/rodolfoviolla/crud-go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:postId", controllers.PostsShow)

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:postId", controllers.PostsUpdate)
	r.DELETE("posts/:postId", controllers.PostsDelete)

	r.Run()
}
