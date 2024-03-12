package main

import (
	"go_htmx/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("views/**/*")

	r.GET("/", controllers.NotesIndex)
	r.POST("/notes", controllers.NotesCreate)

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
