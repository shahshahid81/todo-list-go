package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shahshahid81/todo-list-go/controllers"
	"github.com/shahshahid81/todo-list-go/middlewares"
	"github.com/shahshahid81/todo-list-go/models"
)

func main() {

	db := models.ConnectDataBase()

	r := gin.Default()

	authController := controllers.AuthController{Db: db}
	authGroup := r.Group("/api/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)

	todoController := controllers.TodoController{Db: db}
	todoGroup := r.Group("/api/todo")
	todoGroup.Use(middlewares.JwtAuthMiddleware())
	todoGroup.GET("/", todoController.GetAll)
	todoGroup.POST("/", todoController.Create)

	r.Run(":8080")

}
