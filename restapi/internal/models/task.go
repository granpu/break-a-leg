package models

import "time"

type Task struct {
	ID          int       `json:"id" db:"id"`
	Tittle      string    `json:"tittle" db:"tittle"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created_at" db"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db"updated_at"`
}

type CreateTaskInput struct {
	ID          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTaskInput struct {
	Tittle      *string `json:"tittle"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}
