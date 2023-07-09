package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahshahid81/todo-list-go/models"
	"github.com/shahshahid81/todo-list-go/utils"
	"gorm.io/gorm"
)

type CreatePost struct {
	Title       string `json:"title" binding:"required,min=3,max=30"`
	Description string `json:"description" binding:"required,min=10,max=255"`
}

type UpdatePost struct {
	PostId      uint   `json:"postId" binding:"required"`
	Title       string `json:"title,omitempty" binding:"omitempty,min=3,max=30"`
	Description string `json:"description,omitempty" binding:"omitempty,min=10,max=255"`
}

type TodoController struct {
	Db *gorm.DB
}

func (tc *TodoController) GetAll(c *gin.Context) {

	userId, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured"})
		return
	}

	todos := []models.Todo{}
	err = tc.Db.Where("user_id = ?", userId).Find(&todos).Error

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"todos": todos})
}

func (tc *TodoController) Create(c *gin.Context) {

	userId, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured"})
		return
	}

	var requestBody CreatePost

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		UserId:      userId,
	}

	result := tc.Db.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured while creating todo"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "todo created successfully"})
	}

}

func (tc *TodoController) Update(c *gin.Context) {

	userId, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured"})
		return
	}

	var requestBody UpdatePost

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{UserId: userId}
	updateBody := make(map[string]interface{})

	if requestBody.Description != "" {
		updateBody["Description"] = requestBody.Description
	}

	if requestBody.Title != "" {
		updateBody["Title"] = requestBody.Title
	}

	result := tc.Db.Model(&todo).Where("user_id = ? and id = ?", userId, requestBody.PostId).Updates(updateBody)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "todo updated successfully"})
	}

}
