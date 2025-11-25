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

func getTodoById(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID invalide : format UUID requis"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context) {
	var body CreateTodoDTO
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

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID invalide : format UUID requis"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID invalide : format UUID requis"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			if err := c.BindJSON(&todos[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
