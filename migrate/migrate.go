package main

import (
	"github.com/bhaveshsaini/go-crud/initializers"
	"github.com/bhaveshsaini/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
