package main

type TodoDTO struct {
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}
