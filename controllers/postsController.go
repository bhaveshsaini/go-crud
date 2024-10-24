package controllers

import (
	"github.com/bhaveshsaini/go-crud/initializers"
	"github.com/bhaveshsaini/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return the data
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond to user
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id from URL
	id := c.Param("id")

	// get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// respond to user
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond to user
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// delete
	initializers.DB.Delete(&models.Post{}, id)

	// respond to user
	c.JSON(200, gin.H{
		"post": "post deleted",
	})
}
