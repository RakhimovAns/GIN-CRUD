package main

import (
	"gin/Initializers"
	"gin/models"
	"log"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDB()
}
func main() {
	err := Initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatal("Error with migrating Post")
	}
}
