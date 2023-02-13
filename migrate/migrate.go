package main

import (
	"github.com/rodolfoviolla/crud-go/initializers"
	"github.com/rodolfoviolla/crud-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}