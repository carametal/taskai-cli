package models

import "time"

// Category represents the task category.
type Category string

const (
	CategoryWork     Category = "work"
	CategoryPersonal Category = "personal"
	CategoryLearning Category = "learning"
	CategoryHealth   Category = "health"
	CategoryOther    Category = "other"
)

// Priority represents the task priority.
type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)

// Urgency represents the task urgency.
type Urgency string

const (
	UrgencyImmediate Urgency = "immediate"
	UrgencyShortTerm Urgency = "short-term"
	UrgencyLongTerm  Urgency = "long-term"
)

// Status represents the task status.
type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
	StatusArchived   Status = "archived" // Added for cleanup functionality
)

// Task represents a task in the system.
type Task struct {
	ID            string                 `json:"id"`                       // Unique ID for the task
	Title         string                 `json:"title"`                    // Task title
	Description   string                 `json:"description,omitempty"`    // Optional detailed description
	Category      Category               `json:"category,omitempty"`       // Task category
	Priority      Priority               `json:"priority,omitempty"`       // Task priority
	Urgency       Urgency                `json:"urgency,omitempty"`        // Task urgency
	EstimatedTime int                    `json:"estimated_time,omitempty"` // Estimated time in minutes
	DueDate       time.Time              `json:"due_date,omitempty"`       // Due date and time
	CreatedAt     time.Time              `json:"created_at"`               // Creation timestamp
	UpdatedAt     time.Time              `json:"updated_at"`               // Last update timestamp
	CompletedAt   time.Time              `json:"completed_at,omitempty"`   // Completion timestamp
	Status        Status                 `json:"status"`                   // Task status
	Tags          []string               `json:"tags,omitempty"`           // List of tags
	RelatedTasks  []string               `json:"related_tasks,omitempty"`  // List of related task IDs (simpler than full Task objects for JSON)
	Metadata      map[string]interface{} `json:"metadata,omitempty"`       // Other metadata
}
