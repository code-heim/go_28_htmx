package controllers

import (
	"go_htmx/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := []models.Note{
		{
			Name:    "Bug Squashing",
			Content: "A fun term for fixing errors in code. The term \"bug\" is rumored to have originated from a moth trapped in a Harvard Mark II computer in the 1940s.",
		},
		{
			Name:    "Recursion",
			Content: "A mind-bending concept where a function calls itself. Recursion can solve complex problems but can also be tricky to wrap your head around.",
		},
		{
			Name:    "Coffee is Fuel",
			Content: "Just a stereotype, right? Well, programmers often rely on coffee (or their favorite beverage) to stay focused and energized during long coding sessions.",
		},
	}

	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotesCreate(c *gin.Context) {
	var data FormData
	c.Bind(&data)
	// Simulate a delay
	time.Sleep(2 * time.Second)
	c.HTML(http.StatusOK,
		"notes/note.html",
		gin.H{
			"Name":    data.Name,
			"Content": data.Content,
		})
}
