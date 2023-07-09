package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahshahid81/todo-list-go/models"
	"github.com/shahshahid81/todo-list-go/utils"
	"gorm.io/gorm"
)

type CreateTodo struct {
	Title       string `json:"title" binding:"required,min=3,max=30"`
	Description string `json:"description" binding:"required,min=10,max=255"`
}

type UpdateTodo struct {
	Id          uint   `json:"id" binding:"required"`
	Title       string `json:"title,omitempty" binding:"omitempty,min=3,max=30"`
	Description string `json:"description,omitempty" binding:"omitempty,min=10,max=255"`
}

type DeleteTodo struct {
	Id string `uri:"id" binding:"required"`
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

	var requestBody CreateTodo

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

	var requestBody UpdateTodo

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

	result := tc.Db.Model(&todo).Where("user_id = ? and id = ?", userId, requestBody.Id).Updates(updateBody)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": "todo updated successfully"})
	}

}

func (tc *TodoController) Delete(c *gin.Context) {

	userId, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "an error occured"})
		return
	}
	var requestParams DeleteTodo
	if err := c.ShouldBindUri(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := tc.Db.Model(&models.Todo{}).Delete("user_id = ? and id = ?", userId, requestParams.Id)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": "todo deleted successfully"})
	}

}
