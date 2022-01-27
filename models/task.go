package models

// Task represents a task one may work on.
type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
