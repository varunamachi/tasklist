package todo

import (
	"time"
)

// TaskItem - represents a todo item
type TaskItem struct {
	ID          int       `json:"id" db:"id"`
	Heading     string    `json:"heading" db:"heading"`
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

// TaskList -
type TaskList struct {
	Tasks  []*TaskItem `json:"tasks"`
	lastID int
}

func (tl *TaskList) nextID() int {
	id := tl.lastID
	tl.lastID++
	return id
}

// Add -
func (tl *TaskList) Add(items ...*TaskItem) {
	for _, task := range items {
		task.ID = tl.nextID()
		tl.Tasks = append(tl.Tasks, task)
	}
}

// Remove -
func (tl *TaskList) Remove(id int) {
	index := -1
	for i, task := range tl.Tasks {
		if task.ID == id {
			index = i
			break
		}
	}
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
}

// Iterate -
func (tl *TaskList) Iterate(operation func(*TaskItem)) {
	for _, task := range tl.Tasks {
		operation(task)
	}
}

// NewTaskList -
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make([]*TaskItem, 0, 100),
	}
}
