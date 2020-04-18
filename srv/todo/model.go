package todo

import (
	"time"
)

// TaskItem - represents a todo item
type TaskItem struct {
	ID          int       `json:"id" db:"id"`
	Heading     string    `json:"heading, required" db:"heading"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Created     time.Time `json:"created" db:"created"`
	Deadline    time.Time `json:"deadline" db:"deadline"`
	Modified    time.Time `json:"modified" db:"modified"`
}

// NewTaskItem -
func NewTaskItem(heading, description string, deadline time.Time) *TaskItem {
	return &TaskItem{
		Heading:     heading,
		Description: description,
		Status:      "new",
		Created:     time.Now(),
		Modified:    time.Now(),
		Deadline:    deadline,
	}
}

// DbOp - opcode for bulk operations
type DbOp string

// Update -
var Update DbOp = "update"

// Delete -
var Delete DbOp = "delete"

// Create -
var Create DbOp = "create"

// BulkOp - represents batch operation on data storage
type BulkOp struct {
	Op    DbOp
	Items []*TaskItem
}

// Storage - represents a storage backend for the task list
type Storage interface {
	Name() string
	Init() error
	Add(item *TaskItem) error
	Remove(id int) error
	Update(item *TaskItem) error
	Retrieve(id int) (*TaskItem, error)
	Bulk(op BulkOp) error
	RetrieveAll(offset, limit int) ([]*TaskItem, error)
}

var storage Storage

func SetStorage(st Storage) {
	if storage == nil {
		storage = st
	}
}

func GetStorage() Storage {
	return storage
}
