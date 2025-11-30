package services

import (
	"context"
	"errors"
	"time"
	"warnerb47/todo/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TodoServiceImpl struct {
	todoCollection *mongo.Collection
	ctx            context.Context
}

func New(todoCollection *mongo.Collection, ctx context.Context) TodoService {
	return &TodoServiceImpl{
		todoCollection: todoCollection,
		ctx:            ctx,
	}
}

func (todoService *TodoServiceImpl) GetTodoById(id string) (*models.Todo, error) {
	var todo *models.Todo
	objecId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.D{bson.E{Key: "_id", Value: objecId}}
	err = todoService.todoCollection.FindOne(todoService.ctx, query).Decode(&todo)
	return todo, err
}

func (todoService *TodoServiceImpl) GetTodos() ([]*models.Todo, error) {
	var todos []*models.Todo
	cursor, err := todoService.todoCollection.Find(todoService.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(todoService.ctx) {
		var todo models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(todoService.ctx)
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}
	return todos, nil
}

func (todoService *TodoServiceImpl) CreateTodo(todo *models.Todo) error {
	todo.Date = time.Now().Format(time.RFC3339)
	_, err := todoService.todoCollection.InsertOne(todoService.ctx, todo)
	return err
}

func (todoService *TodoServiceImpl) DeleteTodo(id string) error {
	objecId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objecId}}
	result, _ := todoService.todoCollection.DeleteOne(todoService.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (todoService *TodoServiceImpl) UpdateTodo(id string, todo *models.Todo) error {
	objecId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objecId}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "label", Value: todo.Label},
		bson.E{Key: "checked", Value: todo.Checked},
		bson.E{Key: "date", Value: time.Now().Format(time.RFC3339)},
	}}}
	result, _ := todoService.todoCollection.UpdateOne(todoService.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
