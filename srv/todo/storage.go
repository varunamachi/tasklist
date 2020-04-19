package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

// TaskList -
type TaskList struct {
	Tasks  []*TaskItem `json:"tasks"`
	NextID int         `json:"nextID"`
}

func (tl *TaskList) nextID() int {
	id := tl.NextID
	tl.NextID++
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

// JSONStorage -
type JSONStorage struct {
	file *os.File
	list []*TaskItem
}

// Init -
func (pg *JSONStorage) Init() error {
	var err error
	pg.file, err = os.Create(getStorageDir() + "/task_list.json")
	if err != nil {
		return err
	}
	return nil
}

// Add -
func (pg *JSONStorage) Add(ti *TaskItem) error {
	pg.list = append(pg.list, ti)
	raw, err := json.MarshalIndent(pg.list, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(pg.file, string(raw))
	return err
}

// Remove -
func (pg *JSONStorage) Remove(id int) error {
	return nil
}

// Name -
func (pg *JSONStorage) Name() string {
	return "pg"
}

// Update -
func (pg *JSONStorage) Update(item *TaskItem) error {
	return nil
}

// Retrieve -
func (pg *JSONStorage) Retrieve(id int) error {
	return nil
}

// Bulk -
func (pg *JSONStorage) Bulk(op BulkOp) error {
	return nil
}

// RetrieveAll -
func (pg *JSONStorage) RetrieveAll() []*TaskItem {
	return nil
}

func getStorageDir() string {
	user, err := user.Current()
	if err != nil {
		dir, _ := os.Getwd()
		return dir
	}
	return user.HomeDir
}
