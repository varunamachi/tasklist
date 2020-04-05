package todo

import "time"

// TaskItem - represents a todo item
type TaskItem struct {
	Heading     string
	Description string
	Status      string
	Created     time.Time
	Deadline    time.Time
	Modified    time.Time
}

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
	Tasks []*TaskItem
}

// Add -
func (list *TaskList) Add(item ...*TaskItem) {
	list.Tasks = append(list.Tasks, item...)
}

// NewTaskList -
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make([]*TaskItem, 0, 100),
	}
}
