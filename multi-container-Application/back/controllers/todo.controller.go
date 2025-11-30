package controllers

import (
	"net/http"
	"warnerb47/todo/models"
	"warnerb47/todo/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoService
}

func New(todoService services.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (todoController *TodoController) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := todoController.todoService.GetTodoById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (todoController *TodoController) GetTodos(ctx *gin.Context) {
	todos, err := todoController.todoService.GetTodos()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (todoController *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := todoController.todoService.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (todoController *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := todoController.todoService.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (todoController *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo models.Todo
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := todoController.todoService.UpdateTodo(id, &todo)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (todoController *TodoController) RegisterTodoRoutes(routeGroupe *gin.RouterGroup) {
	todoRoute := routeGroupe.Group("/todos")
	todoRoute.GET("/", todoController.GetTodos)
	todoRoute.GET("/:id", todoController.GetTodoById)
	todoRoute.POST("/", todoController.CreateTodo)
	todoRoute.DELETE("/:id", todoController.DeleteTodo)
	todoRoute.PUT("/:id", todoController.UpdateTodo)
}
