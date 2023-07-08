package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahshahid81/todo-list-go/models"
	"github.com/shahshahid81/todo-list-go/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuthController struct {
	Db *gorm.DB
}

type RegisterInput struct {
	Email           string        `json:"email" binding:"required,email"`
	Password        string        `json:"password" binding:"required,min=8,max=24,eqfield=ConfirmPassword"`
	ConfirmPassword string        `json:"confirmPassword" binding:"required,min=8,max=24"`
	FirstName       string        `json:"firstName" binding:"required,min=2,max=20"`
	LastName        string        `json:"lastName" binding:"required,min=2,max=20"`
	DateOfBirth     utils.IsoDate `json:"dateOfBirth" binding:"required" time_format:"2006-01-02"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}

func (ac *AuthController) Register(c *gin.Context) {

	var requestBody RegisterInput

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = requestBody.Email
	u.Password = requestBody.Password
	u.FirstName = requestBody.FirstName
	u.LastName = requestBody.LastName
	u.DateOfBirth = datatypes.Date(requestBody.DateOfBirth)

	_, err := u.SaveUser(ac.Db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registration success"})
}

func (ac *AuthController) Login(c *gin.Context) {

	var requestBody LoginInput

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(ac.Db, requestBody.Email, requestBody.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
