package main

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
	Date    string `json:"date"`
}

type CreateTodoDTO struct {
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}

var todos = []Todo{
	{ID: uuid.New().String(), Label: "Blue Train", Checked: false, Date: time.Now().Format(time.RFC3339)},
	{ID: uuid.New().String(), Label: "Jeru", Checked: false, Date: time.Now().Format(time.RFC3339)},
	{ID: uuid.New().String(), Label: "Sarah Vaughan and Clifford Brown", Checked: false, Date: time.Now().Format(time.RFC3339)},
}
