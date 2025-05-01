package repository

import "taskai-cli/internal/models"

// TaskRepository defines the interface for task persistence operations.
type TaskRepository interface {
	// AddTask adds a new task to the repository.
	AddTask(task models.Task) error

	// GetTaskByID retrieves a task by its ID.
	// Returns an error (e.g., ErrTaskNotFound) if the task is not found.
	GetTaskByID(id string) (models.Task, error)

	// GetAllTasks retrieves all tasks from the repository.
	// Optionally accepts filter criteria (to be defined later).
	GetAllTasks() ([]models.Task, error)

	// UpdateTask updates an existing task in the repository.
	// Returns an error (e.g., ErrTaskNotFound) if the task is not found.
	UpdateTask(task models.Task) error

	// DeleteTask removes a task from the repository by its ID.
	// Returns an error (e.g., ErrTaskNotFound) if the task is not found.
	DeleteTask(id string) error
}

// Define potential errors (can be expanded later)
type RepositoryError string

func (e RepositoryError) Error() string {
	return string(e)
}

const (
	ErrTaskNotFound      RepositoryError = "task not found"
	ErrTaskAlreadyExists RepositoryError = "task already exists"
	ErrDataStore         RepositoryError = "data store error" // Generic data storage error
)
