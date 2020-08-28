package main

import (
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
	r.GET("/questions", func(c *gin.Context) {

	})
	r.POST("/questions", func(c *gin.Context) {

	})
	r.PUT("/questions/:questionID", func(c *gin.Context) {

	})
	r.POST("/questions/:questionID/answers/:answerID", func(c *gin.Context) {

	})
	r.Run()
}
