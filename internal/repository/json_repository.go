package repository

import (
	"encoding/json"
	"errors"
	"os"
	"taskai-cli/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

// JSONTaskRepository implements TaskRepository using a JSON file.
type JSONTaskRepository struct {
	filePath string
	mu       sync.Mutex // Mutex to handle concurrent access
}

// NewJSONTaskRepository creates a new instance of JSONTaskRepository.
// It ensures the data file exists.
func NewJSONTaskRepository(filePath string) (*JSONTaskRepository, error) {
	repo := &JSONTaskRepository{
		filePath: filePath,
	}
	// Ensure the file exists, create if not
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := repo.saveTasks([]models.Task{}); err != nil {
			return nil, err
		}
	}
	return repo, nil
}

// loadTasks reads tasks from the JSON file.
func (r *JSONTaskRepository) loadTasks() ([]models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		// If file doesn't exist (should have been created by New), return empty list or error
		if os.IsNotExist(err) {
			return []models.Task{}, nil // Return empty slice if file not found after init
		}
		return nil, errors.Join(ErrDataStore, err)
	}

	// Handle empty file case
	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, errors.Join(ErrDataStore, err)
	}
	return tasks, nil
}

// saveTasks writes tasks to the JSON file.
func (r *JSONTaskRepository) saveTasks(tasks []models.Task) error {
	// No lock needed here if called only by methods that already hold the lock.
	// However, adding it makes this function safer for potential future direct use.
	// Let's keep it simple for now and assume callers handle locking.
	// r.mu.Lock()
	// defer r.mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ") // Pretty print JSON
	if err != nil {
		return errors.Join(ErrDataStore, err)
	}

	// WriteFile ensures atomic write (creates temp file, writes, then renames)
	if err := os.WriteFile(r.filePath, data, 0644); err != nil { // Use standard file permissions
		return errors.Join(ErrDataStore, err)
	}
	return nil
}

// AddTask adds a new task.
func (r *JSONTaskRepository) AddTask(task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.loadTasksInternal() // Use internal load without lock
	if err != nil {
		return err
	}

	// Assign ID and timestamps
	if task.ID == "" {
		task.ID = uuid.NewString()
	}
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	if task.Status == "" {
		task.Status = models.StatusTodo // Default status
	}

	// Check if ID already exists (optional, UUID should be unique)
	for _, t := range tasks {
		if t.ID == task.ID {
			return ErrTaskAlreadyExists
		}
	}

	tasks = append(tasks, task)
	return r.saveTasksInternal(tasks) // Use internal save without lock
}

// GetTaskByID retrieves a task by ID.
func (r *JSONTaskRepository) GetTaskByID(id string) (models.Task, error) {
	tasks, err := r.loadTasks() // Public load uses lock
	if err != nil {
		return models.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, ErrTaskNotFound
}

// GetAllTasks retrieves all tasks.
func (r *JSONTaskRepository) GetAllTasks() ([]models.Task, error) {
	return r.loadTasks() // Public load uses lock
}

// UpdateTask updates an existing task.
func (r *JSONTaskRepository) UpdateTask(updatedTask models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.loadTasksInternal()
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			// Preserve CreatedAt, update UpdatedAt
			updatedTask.CreatedAt = task.CreatedAt
			updatedTask.UpdatedAt = time.Now()
			tasks[i] = updatedTask
			found = true
			break
		}
	}

	if !found {
		return ErrTaskNotFound
	}

	return r.saveTasksInternal(tasks)
}

// DeleteTask removes a task by ID.
func (r *JSONTaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.loadTasksInternal()
	if err != nil {
		return err
	}

	initialLen := len(tasks)
	filteredTasks := make([]models.Task, 0, initialLen)
	for _, task := range tasks {
		if task.ID != id {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == initialLen {
		return ErrTaskNotFound // Task with the given ID was not found
	}

	return r.saveTasksInternal(filteredTasks)
}

// --- Internal helper methods (without locking, assuming caller holds lock) ---

// loadTasksInternal reads tasks without acquiring the lock.
func (r *JSONTaskRepository) loadTasksInternal() ([]models.Task, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, errors.Join(ErrDataStore, err)
	}
	if len(data) == 0 {
		return []models.Task{}, nil
	}
	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, errors.Join(ErrDataStore, err)
	}
	return tasks, nil
}

// saveTasksInternal writes tasks without acquiring the lock.
func (r *JSONTaskRepository) saveTasksInternal(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return errors.Join(ErrDataStore, err)
	}
	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return errors.Join(ErrDataStore, err)
	}
	return nil
}
