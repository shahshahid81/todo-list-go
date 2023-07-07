package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shahshahid81/todo-list-go/controllers"
	"github.com/shahshahid81/todo-list-go/models"
)

func main() {

	db := models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api/auth")

	authController := controllers.AuthController{Db: db}
	public.POST("/register", authController.Register)

	r.Run(":8080")

}
