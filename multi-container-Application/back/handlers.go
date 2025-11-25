package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context) {
	var body TodoDTO
	if err := c.BindJSON(&body); err != nil {
		return
	}
	var newTodo Todo
	newTodo.ID = uuid.New().String()
	newTodo.Date = time.Now().Format(time.RFC3339)
	newTodo.Label = body.Label
	newTodo.Checked = body.Checked
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
