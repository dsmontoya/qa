package main

import (
	"qa/handlers"
	"qa/models"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := models.Connect(); err != nil {
		panic(err)
	}
	defer models.Close()
	if err := models.Migrate(); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/questions", handlers.GetQuestions)
	r.POST("/questions", handlers.Ask)
	r.PUT("/posts/:id", handlers.UpdatePost)
	r.POST("/questions/:questionID/answers", handlers.Answer)
	r.Run()
}
