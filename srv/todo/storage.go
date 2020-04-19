package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"
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
func (js *JSONStorage) Init() error {
	var err error
	js.file, err = os.Create(getStorageDir() + "/task_list.json")
	if err != nil {
		return err
	}

	var byt []byte
	byt, err = ioutil.ReadAll(js.file)
	if err != nil {
		return err
	}

	js.list = make([]*TaskItem, 0, 100)
	err = json.Unmarshal(byt, js.list)
	if err != nil {
		return err
	}

	return nil
}

// Add -
func (js *JSONStorage) Add(ti *TaskItem) error {
	js.list = append(js.list, ti)
	raw, err := json.MarshalIndent(js.list, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(js.file, string(raw))
	js.file.Sync()
	return err
}

// Remove -
func (js *JSONStorage) Remove(id int) error {
	var index int
	for i := 0; i < len(js.list); i++ {
		if js.list[i].ID == id {
			index = i
			break
		}
	}
	js.list = append(js.list[:index], js.list[index+1:]...)
	raw, err := json.MarshalIndent(js.list, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(js.file, string(raw))
	js.file.Sync()
	return nil
}

// Name -
func (js *JSONStorage) Name() string {
	return "json"
}

// Update -
func (js *JSONStorage) Update(item *TaskItem) error {
	var listItem *TaskItem
	for i := 0; i < len(js.list); i++ {
		if js.list[i].ID == item.ID {
			listItem = js.list[i]
			break
		}
	}
	if listItem != nil {
		listItem.Created = item.Created
		listItem.Deadline = item.Deadline
		listItem.Description = item.Description
		listItem.Heading = item.Heading
		listItem.Status = item.Status
		listItem.Modified = time.Now() // updating modified time
		return nil
	}
	return fmt.Errorf("Task item with id = %d, not found", item.ID)
}

// Retrieve -
func (js *JSONStorage) Retrieve(id int) (*TaskItem, error) {
	var listItem *TaskItem
	for i := 0; i < len(js.list); i++ {
		if js.list[i].ID == id {
			listItem = js.list[i]
			break
		}
	}
	if listItem != nil {
		return listItem, nil
	}

	return nil, fmt.Errorf("Task item with id = %d, not found", id)
}

// Bulk -
func (js *JSONStorage) Bulk(op BulkOp) error {
	return nil
}

// RetrieveAll -
func (js *JSONStorage) RetrieveAll() []*TaskItem {
	return js.list
}

func getStorageDir() string {
	user, err := user.Current()
	if err != nil {
		dir, _ := os.Getwd()
		return dir
	}
	return user.HomeDir
}
