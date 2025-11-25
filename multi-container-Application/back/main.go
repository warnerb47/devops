package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", healthCheck)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", createTodos)
	router.PUT("/todos/:id", updateTodo)

	router.Run("localhost:8080")
}
