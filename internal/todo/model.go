package todo

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending    Status = "Pending"
	StatusInProgress Status = "In progress"
	StatusCompleted  Status = "Completed"
	StatusArchived   Status = "Archived"
)

type Todo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      Status     `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func NewTodo(title string) *Todo {
	now := time.Now()
	return &Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}

}
