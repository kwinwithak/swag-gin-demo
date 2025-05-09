package models

// todo represents data about a task in the todo list
type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}
