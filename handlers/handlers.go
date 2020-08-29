package handlers

import (
	"net/http"
	"qa/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Answer(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("questionID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	question := &models.Post{ID: uint(questionID)}
	if err := models.First(question); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	answer := &models.Post{}
	if err := c.BindJSON(answer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	answer, err = question.Answer(answer.Content)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, answer)
}

func Ask(c *gin.Context) {
	question := &models.Post{}
	if err := c.BindJSON(question); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := models.Create(question); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, question)
}

func GetQuestions(c *gin.Context) {
	questions, err := models.GetQuestions()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, questions)
}

func UpdatePost(c *gin.Context) {
	post := &models.Post{}
	id := c.Param("id")
	if err := c.BindJSON(post); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := models.Where("id = ?", id).Updates(post).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, post)
}
