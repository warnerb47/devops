package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", healthCheck)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/todos", createTodos)
	router.Run("localhost:8080")
}
