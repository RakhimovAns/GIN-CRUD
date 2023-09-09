package controllers

import (
	"gin/Initializers"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Title}
	result := Initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//Get the Posts
	var posts []models.Post
	Initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
	//Respond with them

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	Initializers.DB.Find(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id of url
	id := c.Param("id")
	//Get the data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Find the post were updating
	var post models.Post
	Initializers.DB.Find(&post, id)
	//Update it
	Initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	//Respond it
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsDelete(c *gin.Context) {
	//Get the id  of url
	id := c.Param("id")
	//Delete it
	Initializers.DB.Delete(&models.Post{}, id)
	//Respond it
	c.Status(200)
}
