package main

import (
	"time"
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
	{ID: "571e5eb7-f253-4d50-b008-937ada01586e", Label: "Blue Train", Checked: false, Date: time.Now().Format(time.RFC3339)},
	{ID: "7bbd5369-fd5a-4ccc-ae8a-a974e420a7a4", Label: "Jeru", Checked: false, Date: time.Now().Format(time.RFC3339)},
	{ID: "3aff95e2-3dbe-4443-860b-40641b38557b", Label: "Sarah Vaughan and Clifford Brown", Checked: false, Date: time.Now().Format(time.RFC3339)},
}
