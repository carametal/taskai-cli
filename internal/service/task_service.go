package service

import (
	"time"
	"taskai-cli/internal/models"
	"taskai-cli/internal/repository"
)

// TaskService defines the interface for task business operations.
type TaskService interface {
	// CreateTask creates a new task with the given title and optional details.
	CreateTask(title string, description string, category models.Category, priority models.Priority) (models.Task, error)

	// GetTaskByID retrieves a task by its ID or partial ID.
	GetTaskByID(id string) (models.Task, error)

	// GetAllTasks retrieves all tasks.
	GetAllTasks() ([]models.Task, error)

	// StartTask marks a task as in-progress.
	StartTask(id string) error
	
	// CompleteTask marks a task as done.
	CompleteTask(id string) error

	// DeleteTask removes a task from the repository.
	DeleteTask(id string) error

	// UpdateTask updates an existing task.
	UpdateTask(task models.Task) error
}

// DefaultTaskService implements TaskService using a TaskRepository.
type DefaultTaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new DefaultTaskService with the given repository.
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &DefaultTaskService{
		repo: repo,
	}
}

// CreateTask implements TaskService.CreateTask
func (s *DefaultTaskService) CreateTask(title string, description string, category models.Category, priority models.Priority) (models.Task, error) {
	now := time.Now()
	// UUID is generated in the repository, but we'll create a task object with basic fields
	task := models.Task{
		Title:       title,
		Description: description,
		Category:    category,
		Priority:    priority,
		Status:      models.StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Add the task to the repository (this will generate the ID)
	err := s.repo.AddTask(task)
	if err != nil {
		return models.Task{}, err
	}

	// Get all tasks and find the one we just created (by title and creation time)
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return models.Task{}, err
	}

	// Find the task we just created (match by title and timestamps)
	for _, t := range tasks {
		if t.Title == title && t.CreatedAt.Equal(now) {
			return t, nil
		}
	}

	return models.Task{}, repository.ErrTaskNotFound
}

// GetTaskByID implements TaskService.GetTaskByID
func (s *DefaultTaskService) GetTaskByID(id string) (models.Task, error) {
	return s.repo.GetTaskByID(id)
}

// GetAllTasks implements TaskService.GetAllTasks
func (s *DefaultTaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

// CompleteTask implements TaskService.CompleteTask
func (s *DefaultTaskService) CompleteTask(id string) error {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return err
	}

	task.Status = models.StatusDone
	task.CompletedAt = time.Now()
	task.UpdatedAt = time.Now()

	return s.repo.UpdateTask(task)
}

// DeleteTask implements TaskService.DeleteTask
func (s *DefaultTaskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

// StartTask implements TaskService.StartTask
func (s *DefaultTaskService) StartTask(id string) error {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return err
	}

	// Only update if the task is not already in progress
	if task.Status == models.StatusInProgress {
		return nil
	}

	task.Status = models.StatusInProgress
	task.UpdatedAt = time.Now()

	return s.repo.UpdateTask(task)
}

// UpdateTask implements TaskService.UpdateTask
func (s *DefaultTaskService) UpdateTask(task models.Task) error {
	// Ensure the updated timestamp is set
	task.UpdatedAt = time.Now()
	return s.repo.UpdateTask(task)
}