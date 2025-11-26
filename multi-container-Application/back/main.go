package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", healthCheck)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", createTodos)
	router.PUT("/todos/:id", updateTodo)

	router.Run("localhost:8080")
}
